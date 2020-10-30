package services

import (
	"go-project/config"
	"go-project/models"
	"go-project/structs"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllUsers Fetch all user data
func GetAllUsers(page, pageSize string) (result structs.Pagination, res structs.Response, err error) {
	var user []models.User
	if err = config.DB.Preload("Role").Scopes(Paginate(page, pageSize)).Find(&user).Error; err != nil {
		res.Message = err.Error()
		return
	}
	config.DB.Model(&models.User{}).Count(&total)
	result.Data = user
	result.Page, _ = strconv.Atoi(page)
	result.PageSize, _ = strconv.Atoi(pageSize)
	result.Total = total
	return
}

//CreateUser ... Insert New data
func CreateUser(user *models.User) (res structs.Response, err error) {
	if err = config.DB.Create(user).Error; err != nil {
		res.Message = err.Error()
		return
	}
	res.Message = "user create success"
	return
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(id string) (user models.User, res structs.Response, err error) {
	if err = config.DB.Preload("Role").Where("id = ?", id).First(&user).Error; err != nil {
		res.Message = err.Error()
		return
	}
	return
}

//CheckUserByID ... Fetch only one user by Id
func CheckUserByID(id string) (user models.User, res structs.Response, err error) {
	if user, res, err = GetUserByID(id); err != nil {
		res.Message = err.Error()
		return
	}
	return
}

//GetUserByKeyword ... Fetch only one user by name
func GetUserByKeyword(page, pageSize, keyword string) (result structs.Pagination, res structs.Response, err error) {
	var user []models.User
	if err = config.DB.Preload("Role").Where("name LIKE ?", "%"+keyword+"%").Or("email LIKE ?", "%"+keyword+"%").Scopes(Paginate(page, pageSize)).Find(&user).Error; err != nil {
		res.Message = err.Error()
		return
	}
	config.DB.Model(&models.User{}).Where("name LIKE ?", "%"+keyword+"%").Or("email LIKE ?", "%"+keyword+"%").Count(&total)
	result.Data = user
	result.Page, _ = strconv.Atoi(page)
	result.PageSize, _ = strconv.Atoi(pageSize)
	result.Total = total
	return
}

//UpdateUser ... Update user
func UpdateUser(user models.User) (res structs.Response, err error) {
	if err = config.DB.Save(&user).Error; err != nil {
		res.Message = err.Error()
		return
	}
	res.Message = "user update success"
	return
}

//DeleteUser ... Delete user
func DeleteUser(user models.User, id string) (res structs.Response, err error) {
	config.DB.Where("id = ?", id).Delete(&user)
	res.Message = "user delete success"
	return
}
