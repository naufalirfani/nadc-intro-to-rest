package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) GetBookByID(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	bookID, err := strconv.ParseInt(param.ByName("bookID"), 10, 64)
	if err != nil {
		log.Printf("[internal][GetBookByID] invalid book_id :%+v\n", err)
		return
	}

	stmt, err := h.DB.Prepare(SQLGetBookByID)
	if err != nil{
		log.Printf("[internal][GetBookByID] invalid prepare statement :%+v\n", err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(bookID)
	var books []Book
	for rows.Next() {
		fmt.Println("row:", rows)
		book := Book{}
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Rate,
			&book.Category,
		)
		if err != nil {
			log.Printf("[internal][GetBookByID] fail to scan :%+v\n", err)
			continue
		}
		books = append(books, book)
	}

	bytes, err := json.Marshal(books)
	if err != nil {
		log.Printf("[internal][GetBookByID] fail to marshal data :%+v\n", err)
		return
	}

	renderJSON(w, bytes, http.StatusOK)

}
