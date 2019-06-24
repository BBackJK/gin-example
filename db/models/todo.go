package models

import (
	"gin-example/lib/common"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model        // default made primary key 'id', 'created_at', 'updated_at', 'deleted_at'
	Title      string `gorm:"not null"`
	Completed  bool   `gorm:"default:0;not null"`
}

func (m *Todo) Definition() common.JSON {
	return common.JSON{
		"id":        m.ID,
		"title":     m.Title,
		"completed": m.Completed,
	}
}
