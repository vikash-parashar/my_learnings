// routes/user_routes.go
package routes

import (
	"my_learnings/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup) {
	// User-related routes
	r.GET("/profile", controllers.UserProfile)
}
