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

	/* [GET] /v1/devices
	List all devices
	Example: curl -X GET http://localhost:8080/v1/devices
	Response: [{"id":"1","name":"test","deviceBrand":"test","createdAt":"2021-07-04T16:00:00Z"}]

	*/

	/* [GET] /v1/devices/:id
	Get a device by id
	Example: curl -X GET http://localhost:8080/v1/devices/1
	Response: {"id":"1","name":"test","deviceBrand":"test","createdAt":"2021-07-04T16:00:00Z"}
	*/

	/* [POST] /v1/devices
	Add a new devices
	Example: curl -X POST http://localhost:8080/v1/devices -d '{"name":"test","deviceBrand":"test"}'
	Response: {"uuid":"1"}
	*/

	/* [DELETE] /v1/devices/:id
	Delete a device by id
	Example: curl -X DELETE http://localhost:8080/v1/devices/1
	Response: {}
	*/

	/* [PUT] /v1/devices/:id
	Update whole device object by id
	Example: curl -X PUT http://localhost:8080/v1/devices/1 -d '{"id":"1","name":"test","deviceBrand":"test","createdAt":"2021-07-04T16:00:00Z"}'
	Response: {"id":"1"}
	*/

	/* [PATCH] /v1/devices/:id
	Update partial device object by id
	Example: curl -X PATCH http://localhost:8080/v1/devices/1 -d '{"name":"test","deviceBrand":"test"}'
	Response: {"id":"1"}
	*/

	/* [GET] /v1/devices/search?q=test
	Search devices by brand
	Example: curl -X GET http://localhost:8080/v1/devices/search?q=test
	Response: [{"id":"1","name":"test","deviceBrand":"test","createdAt":"2021-07-04T16:00:00Z"}]
	*/

	router.GET("", devicesRouter.listDevices)
	router.GET("/:id", devicesRouter.getDevice)
	router.POST("", devicesRouter.addDevice)
	router.DELETE("/:id", devicesRouter.deleteDevice)
	router.PUT("/:id", devicesRouter.replaceDevice)
	router.PATCH("/:id", devicesRouter.patchDevice)
	router.GET("/search", devicesRouter.searchDevices)
}

func (dr *DevicesRouter) searchDevices(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		dr.logger.Printf("Error searching devices: %s", "query param is required")
		c.JSON(400, gin.H{
			"message": "query param is required",
		})
		return
	}

	devices, err := dr.devicesService.SearchDevices(c.Request.Context(), query)
	if err != nil {
		dr.logger.Printf("Error searching devices: %s", err)
		c.JSON(500, gin.H{
			"message": "Error searching devices",
		})
		return
	}

	c.JSON(200, devices)
}

func (dr *DevicesRouter) patchDevice(c *gin.Context) {
	var device *model.PatchDeviceRequest
	err := c.BindJSON(&device)
	if err != nil {
		dr.logger.Printf("Error binding device: %s", err)
		c.JSON(400, gin.H{
			"message": "Error binding device",
		})
		return
	}

	deviceID := c.Param("id")
	if deviceID == "" {
		dr.logger.Printf("Error patching device: %s", err)
		c.JSON(400, gin.H{
			"message": "Error patching device",
		})
		return
	}
	device.ID = deviceID
	_, err = dr.devicesService.PatchDevice(c.Request.Context(), device)
	if err != nil {
		dr.logger.Printf("Error patching device: %s", err)
		c.JSON(500, gin.H{
			"message": "Error patching device",
		})
		return
	}

	c.JSON(200, device)
}

func (dr *DevicesRouter) replaceDevice(c *gin.Context) {
	var device *model.Device
	err := c.BindJSON(&device)
	if err != nil {
		dr.logger.Printf("Error binding device: %s", err)
		c.JSON(400, gin.H{
			"message": "Error binding device",
		})
		return
	}

	device.ID = c.Param("id")
	_, err = dr.devicesService.ReplaceDevice(c.Request.Context(), device)
	if err != nil {
		dr.logger.Printf("Error replacing device: %s", err)
		c.JSON(500, gin.H{
			"message": "Error replacing device",
		})
		return
	}

	c.JSON(200, device)
}

func (dr *DevicesRouter) deleteDevice(c *gin.Context) {
	err := dr.devicesService.DeleteDevice(c.Request.Context(), c.Param("id"))
	if err != nil {
		dr.logger.Printf("Error deleting device: %s", err)
		c.JSON(500, gin.H{
			"message": "Error deleting device",
		})
		return
	}

	c.JSON(200, nil)
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

	c.JSON(200, devices)
}

func (dr *DevicesRouter) getDevice(c *gin.Context) {
	id := c.Param("id")
	device, err := dr.devicesService.GetDevice(c.Request.Context(), id)
	if err != nil {
		dr.logger.Printf("Error getting device: %s", err)
		c.JSON(404, gin.H{
			"message": "Error getting device",
		})
		return
	}

	if device == nil {
		dr.logger.Printf("Device not found: %s", id)
		c.JSON(404, gin.H{
			"message": "Device not found",
		})
		return
	}

	c.JSON(200, device)
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

	c.JSON(201, model.NewDeviceResponse{
		UUID: *id,
	})
}
