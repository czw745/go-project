package models

//Permission struct
type Permission struct {
	Default
	PermissionCategoryID int                `gorm:"column:permission_category_id; index" json:"permission_category_id"`
	PermissionCategory   PermissionCategory `gorm:"foreignKey:PermissionCategoryID" json:"permission_category"`
	Name                 string             `gorm:"column:name" json:"name"`
	DisplayName          string             `gorm:"column:display_name" json:"display_name"`
	GuardName            string             `gorm:"column:guard_name" json:"guard_name"`
}

//TableName ... PermissionCategory
func (b *Permission) TableName() string {
	return "permissions"
}

//PermissionResponse struct
type PermissionResponse struct {
	ID                   uint   `json:"id"`
	Name                 string `json:"name"`
	DisplayName          string `json:"display_name"`
	PermissionCategoryID uint   `gorm:"foreignKey:ID; index" json:"permission_category_id"`
}
