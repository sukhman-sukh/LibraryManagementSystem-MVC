package models

import (
	"database/sql"
	"lib-manager/pkg/types"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Fetch Books From Database
func GetBooks(db *sql.DB) (string, []types.Books) {
	var books []types.Books
	var book types.Books

	rows, _ := db.Query("SELECT bookId, bookName, author, copies FROM books_record")
	defer rows.Close()

	for rows.Next() {
		var bookID, bookName, author string
		var copies int
		if err := rows.Scan(&bookID, &bookName, &author, &copies); err != nil {
			panic(err)
		}
		book.BookId = bookID
		book.BookName = bookName
		book.Author = author
		book.Copies = copies
		books = append(books, book)
	}

	// If the datatype is empty
	if len(books) == 0 {
		book.BookId = "empty"
		book.BookName = "empty"
		book.Author = "empty"
		book.Copies = 0
		books = append(books, book)
	}
	return "OK", books
}

// Fetch List of Books Requested for checkout
func GetRequestBooks(db *sql.DB, userId int, userName string) (string, []types.RequestBooks) {
	var rows *sql.Rows
	var requestBooks []types.RequestBooks
	var requestBook types.RequestBooks

	rows, _ = db.Query("SELECT requests.reqId, requests.bookId ,requests.userId, requests.status ,books_record.bookName  FROM requests INNER JOIN books_record ON requests.bookId = books_record.bookId WHERE requests.userId=? ", userId)
	defer rows.Close()

	for rows.Next() {
		var requestID, bookId, userId, status, bookName string

		if err := rows.Scan(&requestID, &bookId, &userId, &status, &bookName); err != nil {
			panic(err)
		}
		requestBook.RequestId = requestID
		requestBook.BookId = bookId
		requestBook.UserId = userId
		requestBook.Status = status
		requestBook.BookName = bookName
		requestBook.UserName = userName

		requestBooks = append(requestBooks, requestBook)
	}
	// If database is empty
	if len(requestBooks) == 0 {
		requestBook.RequestId = "empty"
		requestBook.BookId = "empty"
		requestBook.UserId = "empty"
		requestBook.Status = "empty"
		requestBook.BookName = "empty"
		requestBook.UserName = "empty"
		requestBooks = append(requestBooks, requestBook)
	}
	return "OK", requestBooks
}

func GetIssuedBooks(db *sql.DB, userId int, admin int, userName string) (string, []types.IssuedBook) {
	var rows *sql.Rows
	var issuedBooks []types.IssuedBook
	var issuedBook types.IssuedBook

	if admin == 1 {
		rows, _ = db.Query("SELECT requests.reqId, requests.bookId ,requests.userId, requests.status ,books_record.bookName  FROM requests INNER JOIN books_record ON requests.bookId = books_record.bookId ")
	}
	defer rows.Close()

	for rows.Next() {
		var requestID, bookId, userId, status, bookName string

		if err := rows.Scan(&requestID, &bookId, &userId, &status, &bookName); err != nil {
			panic(err)
		}
		issuedBook.RequestId = requestID
		issuedBook.BookId = bookId
		issuedBook.UserId = userId
		issuedBook.Status = status
		issuedBook.BookName = bookName
		issuedBook.UserName = userName

		issuedBooks = append(issuedBooks, issuedBook)
	}
	// If database is empty
	if len(issuedBooks) == 0 {
		issuedBook.RequestId = "empty"
		issuedBook.BookId = "empty"
		issuedBook.UserId = "empty"
		issuedBook.Status = "empty"
		issuedBook.BookName = "empty"
		issuedBook.UserName = "empty"
		issuedBooks = append(issuedBooks, issuedBook)
	}
	return "OK", issuedBooks
}

// Fetch list of all clients requesting admin access
func GetAdminRequest(db *sql.DB) (string, []types.AdminRequest) {
	var adminRequests []types.AdminRequest
	var adminRequest types.AdminRequest

	rows, err := db.Query("SELECT adminReq.reqId , adminReq.userId ,adminReq.status , users.userName FROM adminReq INNER JOIN users ON adminReq.userId = users.id")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var requestID, userId, status, userName string
		// var copies int
		if err := rows.Scan(&requestID, &userId, &status, &userName); err != nil {
			panic(err)
		}
		adminRequest.RequestId = requestID
		adminRequest.UserId = userId
		adminRequest.Status = status
		adminRequest.UserName = userName
		adminRequests = append(adminRequests, adminRequest)
	}

	if len(adminRequests) == 0 {

		adminRequest.RequestId = "empty"
		adminRequest.UserId = "empty"
		adminRequest.Status = "empty"
		adminRequest.UserName = "empty"
		adminRequests = append(adminRequests, adminRequest)
	}
	return "OK", adminRequests
}

// Accept admin request
func AdminAccept(writer http.ResponseWriter, request *http.Request, db *sql.DB, requestId string) string {

	rows, err := db.Query("SELECT userId FROM adminReq WHERE reqId = ?", requestId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var userId string
		if err := rows.Scan(&userId); err != nil {
			panic(err)
		}
		db.Exec("UPDATE users SET admin = 1 WHERE id = ? ", userId)
		db.Exec("DELETE FROM adminReq WHERE reqId = ? ", requestId)

	}
	return "OK"
}

// Deny admin request
func AdminDeny(writer http.ResponseWriter, request *http.Request, db *sql.DB, requestId string) string {
	db.Exec("DELETE FROM adminReq WHERE reqId = ? ", requestId)
	return "OK"
}
