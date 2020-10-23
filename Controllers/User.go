package controllers

import (
	"go-project/models"
	"go-project/services"
	"go-project/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var res structs.Pagination
	user, err := services.GetAllUsers()
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user models.User
	var res structs.Pagination
	err := c.BindJSON(&user)
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	err = services.CreateUser(&user)
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
	} else {
		res.Message = "user create success"
		c.JSON(http.StatusOK, res)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	var res structs.Pagination
	id := c.Params.ByName("id")
	user, err := services.GetUserByID(id)
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetUserByName ... Get the user by name
func GetUserByName(c *gin.Context) {
	var res structs.Pagination
	name := c.Query("name")
	user, err := services.GetUserByName(name)
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var res structs.Pagination
	id := c.Params.ByName("id")
	user, err := services.CheckUserByID(id)
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.BindJSON(&user)
	err = services.UpdateUser(user)
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
	} else {
		res.Message = "user update success"
		c.JSON(http.StatusOK, res)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var res structs.Pagination
	id := c.Params.ByName("id")
	user, err := services.CheckUserByID(id)
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	err = services.DeleteUser(user, id)
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
	} else {
		res.Message = "user delete success"
		c.JSON(http.StatusOK, res)
	}
}
