package routes

import (
	"my_learnings/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupAddressRoutes sets up routes for the Address model
func SetupAddressRoutes(r *mux.Router) {
	addressRouter := r.PathPrefix("/addresses").Subrouter()

	addressRouter.HandleFunc("/", controllers.GetAddresses).Methods(http.MethodGet)
	// addressRouter.HandleFunc("/{id}", controllers.GetAddress).Methods(http.MethodGet)
	addressRouter.HandleFunc("/", controllers.CreateAddress).Methods(http.MethodPost)
	// addressRouter.HandleFunc("/{id}", controllers.UpdateAddress).Methods(http.MethodPut)
	// addressRouter.HandleFunc("/{id}", controllers.DeleteAddress).Methods(http.MethodDelete)
}
