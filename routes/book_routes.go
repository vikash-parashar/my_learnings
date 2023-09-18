// routes/book_routes.go
package routes

import (
	"my_learnings/controllers"

	"github.com/gin-gonic/gin"
)

func SetupBookRoutes(r *gin.RouterGroup) {
	book := r.Group("/book")
	{
		book.GET("/", controllers.ListBooks)
		book.GET("/:id", controllers.GetBook)
		book.POST("/", controllers.CreateBook)
		book.PUT("/:id", controllers.UpdateBook)
		book.DELETE("/:id", controllers.DeleteBook)
	}
}
