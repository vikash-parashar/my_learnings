package routes

import (
	"my_learnings/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupBookRoutes sets up routes for the Book model
func SetupBookRoutes(r *mux.Router) {
	bookRouter := r.PathPrefix("/books").Subrouter()

	bookRouter.HandleFunc("/", controllers.GetBooks).Methods(http.MethodGet)
	// bookRouter.HandleFunc("/{id}", controllers.GetBookByID).Methods(http.MethodGet)
	bookRouter.HandleFunc("/", controllers.CreateBook).Methods(http.MethodPost)
	// bookRouter.HandleFunc("/{id}", controllers.UpdateBook).Methods(http.MethodPut)
	// bookRouter.HandleFunc("/{id}", controllers.DeleteBook).Methods(http.MethodDelete)
}
