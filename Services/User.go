package services

import (
	"go-project/config"
	"go-project/models"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllUsers Fetch all user data
func GetAllUsers() (user []models.User, err error) {
	if err = config.DB.Preload("Company").Find(&user).Error; err != nil {
		return
	}
	return
}

//CreateUser ... Insert New data
func CreateUser(user *models.User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(id string) (user models.User, err error) {
	if err = config.DB.Preload("Company").Where("id = ?", id).First(&user).Error; err != nil {
		return
	}
	return
}

//CheckUserByID ... Fetch only one user by Id
func CheckUserByID(id string) (user models.User, err error) {
	if user, err = GetUserByID(id); err != nil {
		return
	}
	return
}

//GetUserByName ... Fetch only one user by name
func GetUserByName(name string) (user []models.User, err error) {
	if err = config.DB.Preload("Company").Where("name LIKE ?", "%"+name+"%").Find(&user).Error; err != nil {
		return
	}
	return
}

//UpdateUser ... Update user
func UpdateUser(user models.User) (err error) {
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user models.User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(&user)
	return nil
}
