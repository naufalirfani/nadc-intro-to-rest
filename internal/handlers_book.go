package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) GetBookByID(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	//ambil param
	bookID := param.ByName("bookID")

	//persiapkan prepare statement sql
	stmt, err := h.DB.Prepare(SQLGetBookByID)
	if err != nil{
		log.Printf("[internal][GetBookByID] invalid prepare statement :%+v\n", err)
		return
	}
	defer stmt.Close()

	//execute query
	rows, err := stmt.Query(bookID)
	var books []Book
	for rows.Next() {
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

	//marshal array books
	bytes, err := json.Marshal(books)
	if err != nil {
		log.Printf("[internal][GetBookByID] fail to marshal data :%+v\n", err)
		return
	}

	renderJSON(w, bytes, http.StatusOK)
}

func (h *Handler) InsertNewBook(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	// baca body
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		renderJSON(w, []byte(`
			message: "Failed to read body"
		`), http.StatusBadRequest)
		return
	}

	// parse body
	var book Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		log.Printf("[internal][InsertNewBook] error unmarshal :%+v\n", err)
		return
	}

	//persiapkan prepare statement sql
	stmt, err := h.DB.Prepare(SQLInsertNewBook)
	if err != nil{
		log.Printf("[internal][InsertNewBook] invalid prepare statement :%+v\n", err)
		return
	}
	defer stmt.Close()

	//execute query
	_, err = stmt.Query(book.Title, book.Author, book.Rate, book.Category)
	if err != nil {
		log.Printf("[internal][InsertNewBook] error inserting new book :%+v\n", err)
		return
	}

	renderJSON(w, []byte(`
	{
		status: "success",
		message: "Insert book success!"
	}
	`), http.StatusOK)
}

func (h *Handler) UpdateBookRating(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	//ambil param
	bookID := param.ByName("bookID")

	//baca body
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		renderJSON(w, []byte(`
			message: "Failed to read body"
		`), http.StatusBadRequest)
		return
	}

	// parse body
	var book Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		log.Printf("[internal][UpdateBookRating] error unmarshal :%+v\n", err)
		return
	}

	//persiapkan prepare statement sql
	stmt, err := h.DB.Prepare(SQLUpdateBookRating)
	if err != nil{
		log.Printf("[internal][UpdateBookRating] invalid prepare statement :%+v\n", err)
		return
	}
	defer stmt.Close()

	//execute query
	_, err = stmt.Query(book.Rate, bookID)
	if err != nil {
		log.Printf("[internal][UpdateBookRating] error updating rating :%+v\n", err)
		return
	}

	renderJSON(w, []byte(`
	{
		status: "success",
		message: "Update book success!"
	}
	`), http.StatusOK)
}

func (h *Handler) DeleteBookByID(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	bookID := param.ByName("bookID")

	//persiapkan prepare statement sql
	stmt, err := h.DB.Prepare(SQLDeleteBookByID)
	if err != nil{
		log.Printf("[internal][DeleteBookByID] invalid prepare statement :%+v\n", err)
		return
	}
	defer stmt.Close()

	//execute query
	_, err = stmt.Query(bookID)
	if err != nil {
		log.Printf("[internal][DeleteBookByID] error inserting new book :%+v\n", err)
		return
	}

	renderJSON(w, []byte(`
	{
		status: "success",
		message: "Delete book success!"
	}
	`), http.StatusOK)
}

func (h* Handler) SearchBookByName(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	// baca body
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		renderJSON(w, []byte(`
			message: "Failed to read body"
		`), http.StatusBadRequest)
		return
	}

	// parse body
	var book Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		log.Printf("[internal][SearchBookByName] error unmarshal :%+v\n", err)
		return
	}

	//persiapkan prepare statement sql
	stmt, err := h.DB.Prepare(SQLSearchBookByName)
	if err != nil{
		log.Printf("[internal][SearchBookByName] invalid prepare statement :%+v\n", err)
		return
	}
	defer stmt.Close()

	//execute query
	rows, err := stmt.Query("%"+book.Title+"%")
	var books []Book
	for rows.Next() {
		book := Book{}
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Rate,
			&book.Category,
		)
		if err != nil {
			log.Printf("[internal][SearchBookByName] fail to scan :%+v\n", err)
			continue
		}
		books = append(books, book)
	}

	//marshal array books
	bytes, err := json.Marshal(books)
	if err != nil {
		log.Printf("[internal][SearchBookByName] fail to marshal data :%+v\n", err)
		return
	}

	renderJSON(w, bytes, http.StatusOK)
}
