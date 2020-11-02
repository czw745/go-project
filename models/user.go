package models

type User struct {
	Default
	Name            string  `gorm:"column:name; NOT NULL" json:"name" binding:"required"`
	Email           string  `gorm:"column:email; NOT NULL" json:"email" binding:"required"`
	EmailVerifiedAt *string `gorm:"column:email_verified_at" json:"email_verified_at"`
	Password        string  `gorm:"column:password; NOT NULL" json:"password" binding:"required"`
	Status          uint    `gorm:"column:status; NOT NULL" json:"status"`
	Deletable       uint    `gorm:"column:deletable; NOT NULL" json:"deletable"`
	Role            []Role  `gorm:"many2many:user_role;" json:"role"`
}

func (b *User) TableName() string {
	return "users"
}