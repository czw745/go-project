package config

import "go-project/models"

//Migrate ... init data
func Migrate() {
	DB.AutoMigrate(&models.Role{}, &models.User{}, &models.PermissionCategory{}, &models.Permission{})
}
