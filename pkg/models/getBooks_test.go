package models

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	// Set up the mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	// Creating a new instance of the HTTP request and response
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()

	expectedRows := sqlmock.NewRows([]string{"bookId", "bookName", "author", "copies"}).
		AddRow("1", "Book 1", "Author 1", 10).
		AddRow("2", "Book 2", "Author 2", 5)

	mock.ExpectQuery("SELECT bookId, bookName, author, copies FROM books_record").WillReturnRows(expectedRows)

	// Calling the function to be tested
	msg, books := GetBooks(res, req)

	// Compute the result
	assert.Equal(t, "OK", msg)
	assert.Len(t, books, 2) // Assuming 2 rows were returned in the test data

	// Ensure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
