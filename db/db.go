package db

import (
	"fmt"
	"os"

	"gin-example/db/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Initialize() (*gorm.DB, error) {
	config := os.Getenv("DB_CONFIG")

	db, err := gorm.Open("mysql", config)
	db.LogMode(true)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connecting to MYSQL")
	models.Migrate(db)

	return db, err
}

// db setting in handler similar jwt
func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
