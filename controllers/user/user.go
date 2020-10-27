package user

import (
	"fmt"
	"go-project/models"
	"go-project/services"
	"go-project/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Get ... Get Users
func Get(c *gin.Context) {
	page := c.Query("page")
	pageSize := c.Query("page_size")
	result, res, err := services.GetAllUsers(page, pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

//Create ... Create User
func Create(c *gin.Context) {
	var user models.User
	var res structs.Response
	err := c.BindJSON(&user)
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res, err = services.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

//Show ... Get the user by id
func Show(c *gin.Context) {
	id := c.Params.ByName("id")
	user, res, err := services.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//Search ... search the user by keyword
func Search(c *gin.Context) {
	keyword := c.Query("keyword")
	page := c.Query("page")
	pageSize := c.Query("page_size")
	result, res, err := services.GetUserByKeyword(page, pageSize, keyword)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

//UpdateUser ... Update the user information
func Update(c *gin.Context) {
	id := c.Params.ByName("id")
	user, res, err := services.CheckUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	fmt.Println("user", user)
	c.BindJSON(&user)
	res, err = services.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// //DeleteUser ... Delete the user
func Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	user, res, err := services.CheckUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res, err = services.DeleteUser(user, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, res)
	}
}
