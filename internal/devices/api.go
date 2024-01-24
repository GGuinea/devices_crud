package devices

import (
	"github.com/gin-gonic/gin"
)

func BuildRoutes(router *gin.RouterGroup) {
	router.GET("", listDevices)
}

func listDevices(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "listing",
	})
}
