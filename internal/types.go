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
	ID       int     `json:"book_id" db:"book_id"`
	Title    string  `json:"book_title" db:"book_name"`
	Author   string  `json:"book_author_name" db:"book_author_name"`
	Category string  `json:"book_category" db:"book_category"`
	Rate     float32 `json:"book_rate" db:"book_category"`
}
