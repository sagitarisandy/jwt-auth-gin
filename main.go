package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sagitarisandy/jwt-auth-gin/initializers"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
