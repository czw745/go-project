package Controllers

import (
	"fmt"
	"go-project/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type respone struct {
	Message string `json:"message"`
	Result  bool   `json:"result"`
}

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUsers(&user)

	if err != nil {
		r := &respone{
			Message: "user create fail",
			Result:  false,
		}
		c.JSON(http.StatusOK, *r)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.CreateUser(&user)
	fmt.Println("message", err)

	if err != nil {
		r := &respone{
			Message: "user create fail",
			Result:  false,
		}
		c.JSON(http.StatusOK, *r)
	} else {
		r := &respone{
			Message: "user create success",
			Result:  true,
		}
		c.JSON(http.StatusOK, *r)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetUserByID(&user, id)
	if err != nil {
		fmt.Println(user, err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&user, id)
	if err != nil {
		r := &respone{
			Message: "cant't find the user, user update fail",
			Result:  false,
		}
		c.JSON(http.StatusOK, *r)
	} else {
		c.BindJSON(&user)
		err = Models.UpdateUser(&user, id)
		if err == nil {
			r := &respone{
				Message: "user update success",
				Result:  true,
			}
			c.JSON(http.StatusOK, *r)
		}
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		r := &respone{
			Message: "cant't find the user, user delete fail",
			Result:  false,
		}
		c.JSON(http.StatusOK, *r)
	} else {
		r := &respone{
			Message: "user delete success",
			Result:  true,
		}
		c.JSON(http.StatusOK, *r)
	}
}
