package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sagitarisandy/jwt-auth-gin/controllers"
	"github.com/sagitarisandy/jwt-auth-gin/initializers"
	"github.com/sagitarisandy/jwt-auth-gin/middleware"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
