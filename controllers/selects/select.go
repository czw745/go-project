package selects

import (
	"go-project/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

//SelectRoles ... select roles
func SelectRoles(c *gin.Context) {
	role, res, err := services.SelectRoles()
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, role)
	}
}
