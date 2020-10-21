package Routes

import (
	"go-project/Controllers"

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
			user.GET("/query", Controllers.GetUsers)
			user.POST("/create", Controllers.CreateUser)
			user.GET("/show/:id", Controllers.GetUserByID)
			user.GET("/search", Controllers.GetUserByName)
			user.PUT("/update/:id", Controllers.UpdateUser)
			user.DELETE("/delete/:id", Controllers.DeleteUser)
		}

	}
	return r
}
