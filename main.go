package main

import (
	"devices_crud/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.NewConfig()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	router.Run(":" + config.Router.Port)
}
