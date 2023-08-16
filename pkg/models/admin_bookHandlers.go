package models

import (
	"database/sql"
	"lib-manager/pkg/types"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// Add Books to Database
func AdminAdd(writer http.ResponseWriter, request *http.Request, db *sql.DB, bookname string, Author string, Copies string) string {

	var book types.Books

	rows, _ := db.Query("SELECT * FROM books_record WHERE bookName = ?", bookname)

	if !(rows.Next()) {
		db.Exec("INSERT INTO books_record (bookName, author, copies) VALUES (?, ? ,?)", bookname, Author, Copies)
		return "OK"
	}

	if err := rows.Scan(&book.BookId, &book.BookName, &book.Author, &book.Copies); err != nil {
		panic(err)
	}
	Copy, _ := strconv.Atoi(Copies)
	Copy = book.Copies + Copy
	db.Exec("UPDATE books_record SET copies = ? where bookName = ?", Copy, bookname)

	return "OK"
}

// Remove books from the database
func AdminRemove(writer http.ResponseWriter, request *http.Request, db *sql.DB, bookId string, Copies string) string {

	db.Exec("UPDATE books_record SET copies = ? where bookId = ?", Copies, bookId)
	return "OK"
}

// Approve checkin of books requested by the user by the admin
func AdminCheckin(writer http.ResponseWriter, request *http.Request, db *sql.DB, requestId string) string {

	rows, err := db.Query("SELECT bookId FROM requests WHERE reqId = ?", requestId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var bookId string
		if err := rows.Scan(&bookId); err != nil {
			panic(err)
		}
		db.Exec("DELETE FROM requests WHERE reqId = ? ", requestId)

		rows, err := db.Query("SELECT copies FROM books_record WHERE bookId = ?", bookId)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			var copies int
			if err := rows.Scan(&copies); err != nil {
				panic(err)
			}
			FinalCopies := copies + 1
			db.Exec("UPDATE books_record SET copies =? where bookId =?", FinalCopies, bookId)
		}
	}
	return "OK"
}

// Approve checkoiut of books requested by the user by the admin
func AdminCheckout(writer http.ResponseWriter, request *http.Request, db *sql.DB, requestId string) string {

	rows, _ := db.Query("SELECT bookId FROM requests WHERE reqId = ?", requestId)
	defer rows.Close()

	for rows.Next() {
		var bookId string
		if err := rows.Scan(&bookId); err != nil {
			panic(err)
		}

		db.Exec("UPDATE requests SET status = 0 WHERE reqId = ? ", requestId)

		rows, _ := db.Query("SELECT copies FROM books_record WHERE bookId = ?", bookId)
		defer rows.Close()

		for rows.Next() {
			var copies int
			if err := rows.Scan(&copies); err != nil {
				panic(err)
			}
			FinalCopies := copies - 1
			db.Exec("UPDATE books_record SET copies =? where bookId =?", FinalCopies, bookId)
		}
	}
	return "OK"
}
