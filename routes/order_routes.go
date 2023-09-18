// routes/order_routes.go
package routes

import (
	"my_learnings/controllers"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(r *gin.RouterGroup) {
	order := r.Group("/order")
	{
		order.GET("/", controllers.ListOrders)
		order.GET("/:id", controllers.GetOrder)
		order.POST("/", controllers.CreateOrder)
		order.PUT("/:id", controllers.UpdateOrder)
		order.DELETE("/:id", controllers.DeleteOrder)
	}
}
