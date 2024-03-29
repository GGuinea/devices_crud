package main

import (
	"devices_crud/config"

	"devices_crud/internal/devices"
	"devices_crud/internal/drivers/graph"
	"log"
	"os"
)

func main() {
	config := config.NewConfig()

	logger := log.New(os.Stdout, "devices_crud: ", log.Ldate|log.Ltime|log.Lshortfile)
	devicesDependencies := devices.NewDevicesDependencies(
		&devices.DeviceDependencies{UseMocks: config.DevicesService.UseMocks, Logger: logger})

	// router := gin.Default()
	// rest.BuildRoutes(router, devicesDependencies)
	// router.Run(":" + config.Router.Port)
	gqlServer := graph.RunServer(devicesDependencies)
	gqlServer.Run()
}
