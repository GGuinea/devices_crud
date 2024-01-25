package devices

import (
	"devices_crud/internal/devices/app"

	"github.com/gin-gonic/gin"
)

type DevicesRouter struct {
	devicesService *app.DeviceService
}

func BuildRoutes(router *gin.RouterGroup, devicesDeps *DependencyTree) {
	devicesRouter := &DevicesRouter{
		devicesService: devicesDeps.DeviceSerivce,
	}

	router.GET("", devicesRouter.listDevices)
}

func (dr *DevicesRouter) listDevices(c *gin.Context) {
	devices, err := dr.devicesService.GetAllDevices(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": devices,
	})
}
