package routes

import (
	"gin-example/routes/test"
	"gin-example/routes/todo"

	"github.com/gin-gonic/gin"
)

// route grouping
func RootRoutes(r *gin.Engine) {
	root := r.Group("/")
	{
		test.TestRoutes(root)
		todo.TodoRoutes(root)
	}
}
