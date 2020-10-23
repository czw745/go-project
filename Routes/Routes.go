package routes

import (
	"go-project/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	grp1 := r.Group("/api/v1")
	{
		user := grp1.Group("/user")
		{
			user.GET("/query", controllers.GetUsers)
			user.POST("/create", controllers.CreateUser)
			user.GET("/show/:id", controllers.GetUserByID)
			user.GET("/search", controllers.GetUserByName)
			user.PUT("/update/:id", controllers.UpdateUser)
			user.DELETE("/delete/:id", controllers.DeleteUser)
		}
	}
	return r
}
