package Controllers

import (
	"fmt"
	"go-project/Models"
	"go-project/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message"`
}

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var user []Models.User
	err := Services.GetAllUsers(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, &response{Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user Models.User
	err := c.BindJSON(&user)
	fmt.Println("err", err)
	if err != nil {
		c.JSON(http.StatusBadRequest, &response{Message: err.Error()})
		return
	}
	err = Services.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, &response{Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, &response{Message: "user create success"})
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Services.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &response{Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetUserByName ... Get the user by name
func GetUserByName(c *gin.Context) {
	name := c.Query("name")
	user, err := Services.GetUserByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, &response{Message: err.Error()})
	} else {
		fmt.Println("user", user)
		c.JSON(http.StatusOK, user)
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Services.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &response{Message: err.Error()})
	} else {
		c.BindJSON(&user)
		err = Services.UpdateUser(&user, id)
		if err == nil {
			c.JSON(http.StatusOK, &response{Message: "user update success"})
		}
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Services.DeleteUser(&user, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &response{Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, &response{Message: "user delete success"})
	}
}
