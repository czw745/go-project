package role

import (
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
	result, res, err := services.GetAllRoles(page, pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

//Create ... Create User
func Create(c *gin.Context) {
	var role models.Role
	var res structs.Response
	err := c.BindJSON(&role)
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res, err = services.CreateRole(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

//Show ... Get the user by id
func Show(c *gin.Context) {
	id := c.Params.ByName("id")
	data, res, err := services.GetRoleByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

//Search ... search the role by keyword
func Search(c *gin.Context) {
	keyword := c.Query("keyword")
	page := c.Query("page")
	pageSize := c.Query("page_size")
	result, res, err := services.GetRoleByKeyword(page, pageSize, keyword)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

//Update ... Update the role information
func Update(c *gin.Context) {
	id := c.Params.ByName("id")
	role, res, err := services.CheckRoleByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	err = c.BindJSON(&role)
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res, err = services.UpdateRole(role)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

//Delete ... Delete the role
func Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	data, res, err := services.CheckRoleByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res, err = services.DeleteRole(data, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, res)
	}
}
