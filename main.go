package main

import (
	"my_learnings/dbconfig"
	"my_learnings/middleware"
	"my_learnings/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	_, err := dbconfig.InitDB()
	if err != nil {
		panic(err)
	}

	// Create a new Gin router
	r := gin.Default()

	// Use the AuthMiddleware for all routes
	r.Use(middleware.AuthMiddleware)

	// Initialize routes
	routes.SetupRoutes(r)

	// Start the server
	r.Run(":8080")
}
