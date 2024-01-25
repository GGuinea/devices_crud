package main

import (
	"devices_crud/config"
	"devices_crud/internal/devices"
	"devices_crud/internal/drivers/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.NewConfig()

	devicesDependencies := devices.NewDevicesDependencies(&devices.DeviceDependencies{UseMocks: config.DevicesService.UseMocks})


	router := gin.Default()
	rest.BuildRoutes(router, devicesDependencies)
	router.Run("localhost:" + config.Router.Port)
}
