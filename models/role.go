package models

type Role struct {
	Default
	Name        string `gorm:"column:name; NOT NULL" json:"name,omitempty" binding:"required"`
	DisplayName string `gorm:"column:display_name; NOT NULL" json:"display_name,omitempty" binding:"required"`
	Status      int    `gorm:"column:status; NOT NULL" json:"status"`
	Deletable   int    `gorm:"column:deletable; NOT NULL" json:"deletable"`
}

func (b *Role) TableName() string {
	return "roles"
}

type RoleSelect struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}
