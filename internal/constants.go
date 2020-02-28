package internal

const(
	SQLGetBookByID = `	SELECT book_id, book_name, book_author_name, book_rate, book_category 
						FROM nadc_mst_book 
						WHERE book_id = $1`
	SQLInsertNewBook = ` INSERT INTO nadc_mst_book (book_name, book_author_name, book_rate, book_category)
						 VALUES ($1, $2, $3, $4)
						 RETURNING book_id`
	SQLUpdateBookRating = `	UPDATE nadc_mst_book SET book_rate = $1
							WHERE book_id = $2`
	SQLDeleteBookByID = `	DELETE FROM nadc_mst_book
							WHERE book_id = $1`
	SQLSearchBookByName = `SELECT book_id, book_name, book_author_name, book_rate, book_category 
							FROM nadc_mst_book 
							WHERE book_name ILIKE $1`
)
