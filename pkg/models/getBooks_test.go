package models


import (
	"fmt"
    "testing"
    "github.com/DATA-DOG/go-sqlmock"
    "github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {

	fmt.Println("Hello, world!")
    // Create a mock database connection
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Error creating mock database connection: %v", err)
    }
    defer db.Close()

    // Define your expected query and rows
    rows := sqlmock.NewRows([]string{"bookId", "bookName", "author", "copies"}).
        AddRow("1", "Book 1", "Author 1", 5).
        AddRow("2", "Book 2", "Author 2", 3)

    // Set up the mock query expectation
    mock.ExpectQuery("SELECT bookId, bookName, author, copies FROM books_record").
        WillReturnRows(rows)

    // Call the function with the mock DB
    status, books := GetBooks(db)
	fmt.Println(status , books)

    // Check the returned status and books
    assert.Equal(t, "OK", status)
    assert.Equal(t, 2, len(books)) // Check the number of books returned

    // Check the details of the first book
    assert.Equal(t, "1", books[0].BookId)
    assert.Equal(t, "Book 1", books[0].BookName)
    assert.Equal(t, "Author 1", books[0].Author)
    assert.Equal(t, 5, books[0].Copies)

    // Check the details of the second book
    assert.Equal(t, "2", books[1].BookId)
    assert.Equal(t, "Book 2", books[1].BookName)
    assert.Equal(t, "Author 2", books[1].Author)
    assert.Equal(t, 3, books[1].Copies)

    // Ensure all expectations were met
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("Unfulfilled expectations: %s", err)
    }
}