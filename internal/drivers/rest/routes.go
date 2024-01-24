package rest

import (
	"devices_crud/internal/devices"

	"github.com/gin-gonic/gin"
)

func BuildRoutes(router *gin.Engine) {
	router.GET("/ping", ping)
	devicesPath := router.Group("/v1/devices")

	devices.BuildRoutes(devicesPath)

}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}
