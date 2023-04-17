package main

import (
	"fmt"
	"movies/controllers"
	"movies/initial"
	"os"

	"github.com/gin-gonic/gin"
)

func init(){
	// load .env
	initial.LoadEnv()
	if os.Getenv("GETMOVIES") == "true" {
		fmt.Println("here")
		os.Setenv("GETMOVIES" , "false")
	}
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