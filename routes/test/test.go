package test

import (
	"github.com/gin-gonic/gin"
)

func TestRoutes(r *gin.RouterGroup) {
	tests := r.Group("/test")
	{
		tests.GET("/", get)
		tests.POST("/", post)
	}
}
