package routes

import (
	"github.com/gorilla/mux"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(r *mux.Router) {
	// Setup routes for different models
	SetupBookRoutes(r)
	SetupAuthorRoutes(r)
	SetupUserRoutes(r)
	SetupOrderRoutes(r)
	SetupAddressRoutes(r)

	// Add routes for other models as needed
}
