package main

import (
	"movies/controllers"
	"movies/initial"

	"github.com/gin-gonic/gin"
)

func init(){
	// load .env
	initial.LoadEnv()
	initial.FillingTheMovieTable()
}

func main() {

	//start the gin server
	server := gin.Default()

	//endpoints
	server.POST("/signup" , controllers.SignUp)
	server.POST("/login" , controllers.Login)
	server.POST("/follow" , controllers.Follow )


	//run the server
	server.Run()
}