// controllers/address_controller.go

package controllers

import (
	"encoding/json"
	"my_learnings/models"
	"net/http"
)

// GetAddresses returns a list of addresses.
func GetAddresses(w http.ResponseWriter, r *http.Request) {
    var addresses []models.Address

    // Retrieve addresses from the database using GORM
    if err := models.GetDB().Find(&addresses).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Serialize addresses to JSON and send the response
    json.NewEncoder(w).Encode(addresses)
}

// CreateAddress creates a new address.
func CreateAddress(w http.ResponseWriter, r *http.Request) {
    var address models.Address

    // Decode the request body into an address object
    if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Create the address in the database using GORM
    if err := models.GetDB().Create(&address).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Serialize the created address to JSON and send the response
    json.NewEncoder(w).Encode(address)
}

// Other controller functions for updating, deleting, and handling specific address operations.
