package main

import (
	"fmt"
	"log"
	"my_learnings/dbconfig"
	"my_learnings/routes"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	_, err := dbconfig.InitDB()
	if err != nil {
		panic(err)
	}
	// db.Debug()
	r := mux.NewRouter()

	// Setup routes for different models using the centralized SetupRoutes function
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
