package services

import (
	"go-project/config"
	"go-project/models"
	"go-project/structs"
)

// SelectRoles ... Select roles
func SelectRoles() (roles []structs.SelectRoles, res structs.Response, err error) {
	var role models.Role
	if err = config.DB.Model(&role).Where("status = ?", "1").Find(&roles).Error; err != nil {
		res.Message = err.Error()
		return
	}
	return
}
