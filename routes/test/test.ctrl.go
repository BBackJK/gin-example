package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func post(c *gin.Context) {
	type Test struct {
		Ping string `form:"ping" json:"ping" xml:"ping" binding:"required"`
		Pong string `form:"pong" json:"pong" xml:"pong" binding:"required"`
	}

	var test Test

	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Bad Data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   test,
	})
}
