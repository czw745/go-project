package services

import (
	"go-project/config"
	"go-project/models"
	"go-project/structs"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var total int64

//GetAllUsers Fetch all user data
func GetAllUsers(page, pageSize string) (res structs.Pagination, err error) {
	var user []models.User

	if err = config.DB.Preload("Company").Scopes(Paginate(page, pageSize)).Find(&user).Error; err != nil {
		return
	}

	config.DB.Model(&models.User{}).Group("name").Count(&total)
	res.Data = user
	res.Page, _ = strconv.Atoi(page)
	res.PageSize, _ = strconv.Atoi(pageSize)
	res.Total = total
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

func Paginate(page, pageSize string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(page)
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(pageSize)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
