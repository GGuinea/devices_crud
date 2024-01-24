package main

import (
	"devices_crud/config"
	"devices_crud/internal/drivers/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.NewConfig()

	router := gin.Default()
	rest.BuildRoutes(router)

	router.Run("localhost:" + config.Router.Port)
}
