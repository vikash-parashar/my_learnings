// controllers/author_controller.go

package controllers

import (
	"encoding/json"
	"my_learnings/models"
	"net/http"
)

// GetAuthors returns a list of authors.
func GetAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author

	// Retrieve authors from the database using GORM
	if err := models.GetDB().Find(&authors).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serialize authors to JSON and send the response
	json.NewEncoder(w).Encode(authors)
}

// CreateAuthor creates a new author.
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	// Decode the request body into an author object
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the author in the database using GORM
	if err := models.GetDB().Create(&author).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serialize the created author to JSON and send the response
	json.NewEncoder(w).Encode(author)
}

// Other controller functions for updating, deleting, and handling specific author operations.
