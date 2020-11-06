package models

import (
	"time"

	"gorm.io/gorm"
)

//Default struct
type Default struct {
	ID        uint            `gorm:"column:id" json:"id"`
	CreatedAt time.Time       `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at; index" json:"deleted_at"`
}
