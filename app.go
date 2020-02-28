package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/ridwanakf/nadc-intro-to-rest/internal"
)

func initFlags(args *internal.Args) {
	port := os.Getenv("PORT")
	if port == "" {
		port = *(flag.String("PORT", "5000", "port number for your apps"))
	}
	args.Port = port
}

func initHandler(handler *internal.Handler) error {

	// Initialize SQL DB
	db, err := sql.Open("postgres", "postgres://cgvrypmrzbfzqd:e7869a8aff75aa1cb3747db0d40f9519f1eb6afce2b188050da8fabf21793e68@18.210.51.239:5432/dcffscbsa1dabf?sslmode=require")
	if err != nil {
		return err
	}

	// Check if db connected
	if err = db.PingContext(context.Background()); err != nil {
		return err
	}

	handler.DB = db

	return nil
}

func initRouter(router *httprouter.Router, handler *internal.Handler) {

	router.GET("/", handler.Index)

	// `httprouter` library uses `ServeHTTP` method for it's 404 pages
	router.NotFound = handler
}

func main() {
	args := new(internal.Args)
	initFlags(args)

	handler := new(internal.Handler)
	if err := initHandler(handler); err != nil {
		log.Println("Failed to init handler", err)
		panic(err)
	}

	router := httprouter.New()
	initRouter(router, handler)

	fmt.Println("Apps served on :" + args.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":" + args.Port), router))
}
