package rest

import "github.com/gin-gonic/gin"

func BuildRoutes(router *gin.Engine) {
	router.GET("/ping", ping)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}
