package routes

import (
	"my_learnings/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupAuthorRoutes sets up routes for the Author model
func SetupAuthorRoutes(r *mux.Router) {
	authorRouter := r.PathPrefix("/authors").Subrouter()

	authorRouter.HandleFunc("/", controllers.GetAuthors).Methods(http.MethodGet)
	// authorRouter.HandleFunc("/{id}", controllers.GetAuthorByID).Methods(http.MethodGet)
	authorRouter.HandleFunc("/", controllers.CreateAuthor).Methods(http.MethodPost)
	// authorRouter.HandleFunc("/{id}", controllers.UpdateAuthor).Methods(http.MethodPut)
	// authorRouter.HandleFunc("/{id}", controllers.DeleteAuthor).Methods(http.MethodDelete)
}
