package routes

import (
	"go-project/controllers/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	apiv1 := r.Group("/api/v1")
	{
		apiUser := apiv1.Group("/user")
		{
			apiUser.GET("/query", user.Get)
			apiUser.POST("/create", user.Create)
			apiUser.GET("/show/:id", user.Show)
			apiUser.GET("/search", user.Search)
			apiUser.PUT("/update/:id", user.Update)
			apiUser.DELETE("/delete/:id", user.Delete)
		}
	}
	return r
}
