package routes

import (
	"my_learnings/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupOrderRoutes sets up routes for the Order model
func SetupOrderRoutes(r *mux.Router) {
	orderRouter := r.PathPrefix("/orders").Subrouter()

	orderRouter.HandleFunc("/", controllers.GetOrders).Methods(http.MethodGet)
	// orderRouter.HandleFunc("/{id}", controllers.GetOrderByID).Methods(http.MethodGet)
	orderRouter.HandleFunc("/", controllers.CreateOrder).Methods(http.MethodPost)
	// orderRouter.HandleFunc("/{id}", controllers.UpdateOrder).Methods(http.MethodPut)
	// orderRouter.HandleFunc("/{id}", controllers.DeleteOrder).Methods(http.MethodDelete)
}
