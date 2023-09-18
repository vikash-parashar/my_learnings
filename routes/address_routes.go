// routes/Address_routes.go
package routes

import (
	"my_learnings/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAddressRoutes(r *gin.RouterGroup) {
	address := r.Group("/address")
	{
		address.GET("/", controllers.ListAddresss)
		address.GET("/:id", controllers.GetAddress)
		address.POST("/", controllers.CreateAddress)
		address.PUT("/:id", controllers.UpdateAddress)
		address.DELETE("/:id", controllers.DeleteAddress)
	}
}
