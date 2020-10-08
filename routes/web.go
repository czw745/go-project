package routes

import (
	"go-project/app/Controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.GET("/", Controllers.Ping)
	return r
}
