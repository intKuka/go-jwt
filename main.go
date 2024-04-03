package main

import (
	"jwt-project/controllers"
	"jwt-project/initializers"
	"jwt-project/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	// see users
	r.GET("/users", controllers.GetUsers)

	// add user
	r.POST("/users", controllers.AddUser)

	// validate auth
	r.GET("/auth/test", middleware.RequireAuth, controllers.ValidateAuth)

	// get access token
	r.POST("/auth/token", controllers.GetToken)

	// refresh access token
	r.POST("/auth/refresh", controllers.RefreshToken)

	r.Run()
}
