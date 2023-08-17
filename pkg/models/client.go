package models

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Checkout(writer http.ResponseWriter, request *http.Request, db *sql.DB, bookId string, userID int) error {
	rows, err := db.Query("SELECT copies FROM books_record WHERE bookId = ?", bookId)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var copies int
		if err := rows.Scan(&copies); err != nil {
			return err
		}
		db.Exec("UPDATE books_record SET copies = copies-1 WHERE bookId = ?", bookId)
	}
	db.Exec("INSERT INTO requests (bookId, userId , status) VALUES(?, ? ,?)", bookId, userID, -1)
	return nil
}

func Checkin(writer http.ResponseWriter, request *http.Request, db *sql.DB, requestId string) error {
	rows, err := db.Query("SELECT bookId FROM requests WHERE reqId = ?", requestId)
	if err!= nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var bookId string
		if err := rows.Scan(&bookId); err != nil {
			return err
		}
		db.Exec("DELETE FROM requests WHERE reqId = ? ", requestId)

		rows, err := db.Query("SELECT copies FROM books_record WHERE bookId = ?", bookId)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var copies int
			if err := rows.Scan(&copies); err != nil {
				return err
			}
			FinalCopies := copies + 1
			db.Exec("UPDATE books_record SET copies =? where bookId =?", FinalCopies, bookId)
		}
	}
	return nil
}
