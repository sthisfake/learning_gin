package main

import (
	"movies/controllers"
	"movies/initial"

	"github.com/gin-gonic/gin"
)

func init(){
	// load .env
	initial.LoadEnv()
}

func main() {

	//start the gin server
	server := gin.Default()

	//endpoints
	server.POST("/signup" , controllers.SignUp)
	server.POST("/login" , controllers.Login)


	//run the server
	server.Run()
}