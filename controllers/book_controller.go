// controllers/book_controller.go

package controllers

import (
	"encoding/json"
	"my_learnings/models"
	"net/http"
)

// GetBooks returns a list of books.
func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	// Retrieve books from the database using GORM
	if err := models.GetDB().Find(&books).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serialize books to JSON and send the response
	json.NewEncoder(w).Encode(books)
}

// CreateBook creates a new book.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	// Decode the request body into a book object
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the book in the database using GORM
	if err := models.GetDB().Create(&book).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serialize the created book to JSON and send the response
	json.NewEncoder(w).Encode(book)
}

// Other controller functions for updating, deleting, and handling specific book operations.
