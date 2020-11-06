package models

//Role struct
type Role struct {
	Default
	Name        string       `gorm:"column:name; NOT NULL" json:"name,omitempty" binding:"required"`
	DisplayName string       `gorm:"column:display_name; NOT NULL" json:"display_name,omitempty" binding:"required"`
	Status      int          `gorm:"column:status; NOT NULL" json:"status"`
	Deletable   int          `gorm:"column:deletable; NOT NULL" json:"deletable"`
	Permissions []Permission `gorm:"many2many:role_has_permission;" json:"permissions"`
}

//TableName ... PermissionCategory
func (b *Role) TableName() string {
	return "roles"
}

//RoleResponse struct
type RoleResponse struct {
	ID          uint          `json:"id"`
	Name        string        `json:"name"`
	DisplayName string        `json:"display_name"`
	Status      int           `json:"status"`
	Deletable   int           `json:"deletable"`
	Permissions *[]Permission `gorm:"foreignKey:ID" json:"permissions"`
}
