package Services

import (
	"go-project/Config"
	"go-project/Models"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]Models.User) (err error) {
	if err = Config.DB.Preload("Company").Find(user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *Models.User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *Models.User, id string) (err error) {
	if err = Config.DB.Preload("Company").Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByName ... Fetch only one user by name
func GetUserByName(name string) (user []Models.User, err error) {
	if err = Config.DB.Preload("Company").Where("name LIKE ?", "%"+name+"%").Find(&user).Error; err != nil {
		return
	}
	return
}

//UpdateUser ... Update user
func UpdateUser(user *Models.User, id string) (err error) {
	Config.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *Models.User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Find(user).Error; err != nil {
		return err
	}
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
