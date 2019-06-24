package todo

import (
	"fmt"
	"net/http"

	"gin-example/db/models"
	"gin-example/lib/common"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Todo = models.Todo

func get(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	todoID := c.Param("id")
	var todo Todo

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, common.JSON{
			"status":  http.StatusNotFound,
			"message": "Not Found Data",
		})
		return
	}

	c.JSON(http.StatusOK, common.JSON{
		"status": http.StatusOK,
		"data":   todo.Definition(),
	})
}

func post(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	fmt.Println(db)

	type postBody struct {
		Title string `json:"title" binding:"required"`
	}

	var body postBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, common.JSON{
			"error":   err.Error(),
			"message": "Bad Data",
		})
		return
	}

	todo := Todo{
		Title: body.Title,
	}

	db.NewRecord(todo)
	db.Create(&todo)

	c.JSON(http.StatusOK, common.JSON{
		"data": todo.Definition(),
	})
}

func put(c *gin.Context) {
	var todo Todo

	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	type putBody struct {
		Title     string `json:"title" binding:"required"`
		Completed bool   `json:"completed" binding:"required"`
	}

	var body putBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, common.JSON{
			"error":   err.Error(),
			"message": "Bad Data",
		})
		return
	}

	//checking by id
	db.First(&todo, id)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, common.JSON{
			"status":  http.StatusNotFound,
			"message": "Not Found Data",
		})
		return
	}

	todo.Title = body.Title
	todo.Completed = body.Completed

	db.Save(&todo)

	c.JSON(http.StatusOK, common.JSON{
		"status": http.StatusOK,
		"data":   todo.Definition(),
	})
}

func delete(c *gin.Context) {
	var todo Todo

	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	db.First(&todo, id)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, common.JSON{
			"status":  http.StatusNotFound,
			"message": "Not Found Data",
		})
		return
	}

	db.Delete(&todo)

	c.JSON(http.StatusOK, common.JSON{
		"status":  http.StatusOK,
		"message": "Todo deleted successfully",
	})
}
