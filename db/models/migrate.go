package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Migrate automigrates models using ORM
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Todo{})
	// set up foreign keys
	fmt.Println("Auto Migration has beed processed")
}
