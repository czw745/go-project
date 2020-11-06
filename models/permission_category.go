package models

//PermissionCategory struct
type PermissionCategory struct {
	Default
	Name        string `gorm:"column:name" json:"name"`
	DisplayName string `gorm:"column:display_name" json:"display_name"`
	ParentID    *uint  `gorm:"column:parent_id" json:"parent_id"`
}

//TableName ... PermissionCategory
func (b *PermissionCategory) TableName() string {
	return "permission_category"
}

//PermissionCategoryParentResponse struct
type PermissionCategoryParentResponse struct {
	ID          uint          `json:"id"`
	Name        string        `json:"name"`
	DisplayName string        `json:"display_name"`
	Child       []interface{} `gorm:"type:text" json:"child"`
}

//PermissionCategoryChildResponse struct
type PermissionCategoryChildResponse struct {
	ID          uint                 `json:"id"`
	Name        string               `json:"name"`
	DisplayName string               `json:"display_name"`
	ParentID    uint                 `json:"parent_id"`
	Permissions []PermissionResponse `gorm:"foreignKey:ID" json:"permissions"`
}
