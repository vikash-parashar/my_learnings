// routes/routes.go
package routes

import (
	"my_learnings/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Define routes and path grouping
	// Home Page
	r.GET("/", controllers.HomePage)

	// Registration Group
	registration := r.Group("/register")
	{
		registration.GET("/", controllers.RegistrationPage)
		registration.POST("/", controllers.Registration)
	}

	// Login Group
	login := r.Group("/login")
	{
		login.GET("/", controllers.LoginPage)
		login.POST("/", controllers.Login)
	}

	// Middleware to check if the user is authenticated
	// r.Use(handlers.AuthMiddleware)

	// Grouping for authenticated users
	user := r.Group("/user")
	{
		// Use SetupUserRoutes to define user-specific routes
		SetupUserRoutes(user)
	}

	// Grouping for resources
	SetupBookRoutes(user)    // Use the "user" group for book-related routes
	SetupAuthorRoutes(user)  // Use the "user" group for author-related routes
	SetupOrderRoutes(user)   // Use the "user" group for order-related routes
	SetupAddressRoutes(user) // Use the "user" group for address-related routes
}
