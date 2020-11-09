package routes

import (
	"go-project/controllers/auth"
	"go-project/controllers/permission"
	"go-project/controllers/role"
	"go-project/controllers/selects"
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
		apiAuth := apiv1.Group("/auth")
		{
			apiAuth.POST("/login", auth.Login)
			apiAuth.POST("/logout", auth.Logout)
			apiAuth.GET("/info", auth.Info)
			apiAuth.POST("/refresh-token", auth.RefreshToken)
		}

		apiUser := apiv1.Group("/user")
		{
			apiUser.GET("/query", user.Get)
			apiUser.POST("/create", user.Create)
			apiUser.GET("/show/:id", user.Show)
			apiUser.GET("/search", user.Search)
			apiUser.PUT("/update/:id", user.Update)
			apiUser.DELETE("/delete/:id", user.Delete)
		}

		apiRole := apiv1.Group("/role")
		{
			apiRole.GET("/query", role.Get)
			apiRole.POST("/create", role.Create)
			apiRole.GET("/show/:id", role.Show)
			apiRole.GET("/search", role.Search)
			apiRole.PUT("/update/:id", role.Update)
			apiRole.DELETE("/delete/:id", role.Delete)
		}

		apiPermission := apiv1.Group("/permission")
		{
			apiPermission.GET("/list", permission.Get)
		}

		apiSelect := apiv1.Group("/select")
		{
			apiSelect.GET("/roles", selects.SelectRoles)
		}
	}
	return r
}
