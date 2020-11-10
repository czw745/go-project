package auth

import (
	"fmt"
	"go-project/middleware"
	"go-project/services"
	"go-project/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Login ...
func Login(c *gin.Context) {
	var info structs.AuthLoginRequest
	err := c.BindJSON(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// ip := c.ClientIP()
	// fmt.Println(info)
	result, res, err := services.AuthLogin(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

//Info ... get user info
func Info(c *gin.Context) {
	tokenAuth, err := middleware.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	userID, err := middleware.FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	result, res, err := services.AuthInfo(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

//Logout ...
func Logout(c *gin.Context) {
	// result, res, err := services.AuthLogout(&info)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, res)
	// } else {
	// 	c.JSON(http.StatusOK, result)
	// }
}

//RefreshToken ... refresh token
func RefreshToken(c *gin.Context) {
	refreshToken := c.Request.Header.Get("Authorization")
	fmt.Println(refreshToken)
}
