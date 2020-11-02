package services

import (
	"go-project/config"
	"go-project/models"
	"go-project/structs"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllRoles Fetch all role data
func GetAllRoles(page, pageSize string) (result structs.Pagination, res structs.Response, err error) {
	var role []models.Role
	if err = config.DB.Scopes(Paginate(page, pageSize)).Find(&role).Error; err != nil {
		res.Message = err.Error()
		return
	}
	config.DB.Model(&models.Role{}).Count(&total)
	result.Data = role
	result.Page, _ = strconv.Atoi(page)
	result.PageSize, _ = strconv.Atoi(pageSize)
	result.Total = total
	return
}

//CreateRole ... Insert New data
func CreateRole(role *models.Role) (res structs.Response, err error) {
	if err = config.DB.Create(role).Error; err != nil {
		res.Message = err.Error()
		return
	}
	res.Message = "role create success"
	return
}

//GetRoleByID ... Fetch only one role by Id
func GetRoleByID(id string) (role models.Role, res structs.Response, err error) {
	if err = config.DB.Where("id = ?", id).First(&role).Error; err != nil {
		res.Message = err.Error()
		return
	}
	return
}

//CheckRoleByID ... Fetch only one user by Id
func CheckRoleByID(id string) (role models.Role, res structs.Response, err error) {
	if role, res, err = GetRoleByID(id); err != nil {
		res.Message = err.Error()
		return
	}
	return
}

//GetRoleByKeyword ... Fetch only one role by keyword
func GetRoleByKeyword(page, pageSize, keyword string) (result structs.Pagination, res structs.Response, err error) {
	var role []models.Role
	if err = config.DB.Where("name LIKE ?", "%"+keyword+"%").Or("display_name LIKE ?", "%"+keyword+"%").Scopes(Paginate(page, pageSize)).Find(&role).Error; err != nil {
		res.Message = err.Error()
		return
	}
	config.DB.Model(&models.Role{}).Where("name LIKE ?", "%"+keyword+"%").Or("display_name LIKE ?", "%"+keyword+"%").Count(&total)
	result.Data = role
	result.Page, _ = strconv.Atoi(page)
	result.PageSize, _ = strconv.Atoi(pageSize)
	result.Total = total
	return
}

//UpdateRole ... Update role
func UpdateRole(role models.Role) (res structs.Response, err error) {
	if err = config.DB.Save(&role).Error; err != nil {
		res.Message = err.Error()
		return
	}
	res.Message = "role update success"
	return
}

//DeleteRole ... Delete role
func DeleteRole(role models.Role, id string) (res structs.Response, err error) {
	config.DB.Where("id = ?", id).Delete(&role)
	res.Message = "role delete success"
	return
}

func RoleSelect() (role []models.RoleSelect, res structs.Response, err error) {
	if err = config.DB.Table("roles").Where("status = ?", "1").Where("name != ?", "Super Admin").Find(&role).Error; err != nil {
		res.Message = err.Error()
		return
	}
	return
}
