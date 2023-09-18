package routes

import (
	"my_learnings/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupUserRoutes sets up routes for the User model
func SetupUserRoutes(r *mux.Router) {
	userRouter := r.PathPrefix("/users").Subrouter()

	// userRouter.HandleFunc("/", controllers.GetUsers).Methods(http.MethodGet)
	// userRouter.HandleFunc("/{id}", controllers.GetUserByID).Methods(http.MethodGet)
	userRouter.HandleFunc("/", controllers.CreateUser).Methods(http.MethodPost)
	// userRouter.HandleFunc("/{id}", controllers.UpdateUser).Methods(http.MethodPut)
	// userRouter.HandleFunc("/{id}", controllers.DeleteUser).Methods(http.MethodDelete)
}
