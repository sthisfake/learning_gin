package main

import (
	"jwt/controllers"
	initializers "jwt/initializer"
	"jwt/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDataBase()
}

func main() {

	server := gin.Default()
	server.POST("/signup" , controllers.SignUp)
	server.POST("/login" , controllers.Login)
	server.GET("/validate" , middleware.RequireAuth  , controllers.Validate)

	server.Run()
}