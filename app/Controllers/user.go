package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
