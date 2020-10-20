package Models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint           `gorm:"column:id" json:"id"`
	Name      string         `gorm:"column:name; NOT NULL" json:"name" binding:"required"`
	Nickname  *string        `gorm:"column:nickname" json:"nickname"`
	Email     string         `gorm:"column:email; NOT NULL" json:"email" binding:"required"`
	Password  string         `gorm:"column:password; NOT NULL" json:"password" binding:"required"`
	Phone     *string        `gorm:"column:phone" json:"phone"`
	CreatedAt time.Time      `gorm:"column: created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column: updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (b *User) TableName() string {
	return "users"
}
