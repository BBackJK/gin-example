package todo

import (
	"github.com/gin-gonic/gin"
)

func TodoRoutes(r *gin.RouterGroup) {
	todos := r.Group("/todo")
	{
		todos.GET("/:id", get)
		todos.POST("/", post)
		todos.PUT("/:id", put)
		todos.DELETE("/id", delete)
	}
}
