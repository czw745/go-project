package permission

import (
	"go-project/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Get ... Get Permission List
func Get(c *gin.Context) {
	result, res, err := services.GetParents()
	if err != nil {
		c.JSON(http.StatusBadRequest, res)
	} else {
		c.JSON(http.StatusOK, result)
	}
}
