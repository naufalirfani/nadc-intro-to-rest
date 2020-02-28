package internal

const(
	SQLGetBookByID = `	SELECT * FROM nadc_mst_book 
						WHERE book_id = $1`
)
