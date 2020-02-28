package internal

import (
	"database/sql"
)

// Args used for this application
type Args struct {
	// Port used by this service
	Port string
}

// Handler object used to handle the HTTP API
type Handler struct {
	// DB object that'll be used
	DB *sql.DB
}

// Book struct for database query
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Stock  int64  `json:"stock"`
}
