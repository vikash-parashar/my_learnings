// controllers/order_controller.go

package controllers

import (
	"encoding/json"
	"my_learnings/models"
	"net/http"
)

// GetOrders returns a list of orders.
func GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order

	// Retrieve orders from the database using GORM
	if err := models.GetDB().Find(&orders).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serialize orders to JSON and send the response
	json.NewEncoder(w).Encode(orders)
}

// CreateOrder creates a new order.
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order

	// Decode the request body into an order object
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the order in the database using GORM
	if err := models.GetDB().Create(&order).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serialize the created order to JSON and send the response
	json.NewEncoder(w).Encode(order)
}

// Other controller functions for updating, deleting, and handling specific order operations.
