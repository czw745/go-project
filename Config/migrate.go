package config

import "go-project/models"

func Migrate() {
	DB.AutoMigrate(&models.Role{}, &models.User{}, &models.PermissionCategory{}, &models.Permission{})
}
