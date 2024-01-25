package devices

import (
	"devices_crud/internal/devices/app"
	"devices_crud/internal/devices/model"
	"log"

	"github.com/gin-gonic/gin"
)

type DevicesRouter struct {
	devicesService *app.DeviceService
	logger         *log.Logger
}

func BuildRoutes(router *gin.RouterGroup, devicesDeps *DependencyTree) {
	devicesRouter := &DevicesRouter{
		devicesService: devicesDeps.DeviceSerivce,
		logger:         devicesDeps.Logger,
	}

	router.GET("", devicesRouter.listDevices)
	router.GET("/:id", devicesRouter.getDevice)
	router.POST("", devicesRouter.addDevice)
}

func (dr *DevicesRouter) listDevices(c *gin.Context) {
	devices, err := dr.devicesService.GetAllDevices(c.Request.Context())
	if err != nil {
		dr.logger.Printf("Error getting devices: %s", err)
		c.JSON(500, gin.H{
			"message": "Error getting devices",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": devices,
	})
}

func (dr *DevicesRouter) getDevice(c *gin.Context) {
	device, err := dr.devicesService.GetDevice(c.Request.Context(), c.Param("id"))
	if err != nil {
		dr.logger.Printf("Error getting device: %s", err)
		c.JSON(500, gin.H{
			"message": "Error getting device",
		})
		return
	}

	if device == nil {
		dr.logger.Printf("Device not found: %s", c.Param("id"))
	}

	c.JSON(200, gin.H{
		"data": device,
	})
}

func (dr *DevicesRouter) addDevice(c *gin.Context) {
	var device *model.NewDeviceRequest
	err := c.BindJSON(&device)
	if err != nil {
		dr.logger.Printf("Error binding device: %s", err)
		c.JSON(400, gin.H{
			"message": "Error binding device",
		})
		return
	}

	id, err := dr.devicesService.AddDevice(c.Request.Context(), device)
	if err != nil {
		dr.logger.Printf("Error adding device: %s", err)
		c.JSON(500, gin.H{
			"message": "Error adding device",
		})
		return
	}

	c.JSON(201, gin.H{
		"uuid": id,
	})
}
