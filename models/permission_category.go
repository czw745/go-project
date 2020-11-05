package models

type PermissionCategory struct {
	Default
	Name        string `gorm:"column:name" json:"name"`
	DisplayName string `gorm:"column:display_name" json:"display_name"`
	ParentID    *uint  `gorm:"column:parent_id" json:"parent_id"`
}

func (b *PermissionCategory) TableName() string {
	return "permission_category"
}
