// routes/author_routes.go
package routes

import (
	"my_learnings/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAuthorRoutes(r *gin.RouterGroup) {
	author := r.Group("/author")
	{
		author.GET("/", controllers.ListAuthors)
		author.GET("/:id", controllers.GetAuthor)
		author.POST("/", controllers.CreateAuthor)
		author.PUT("/:id", controllers.UpdateAuthor)
		author.DELETE("/:id", controllers.DeleteAuthor)
	}
}
