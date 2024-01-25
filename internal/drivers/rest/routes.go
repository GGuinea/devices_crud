package rest

import (
	"devices_crud/internal/devices"

	"github.com/gin-gonic/gin"
)

func BuildRoutes(router *gin.Engine, devicesDeps *devices.DependencyTree) {
	router.GET("/ping", ping)
	devicesPath := router.Group("/v1/devices")

	devices.BuildRoutes(devicesPath, devicesDeps)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}
