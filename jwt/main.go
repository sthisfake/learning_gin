package main

import (
	initializers "jwt/initializer"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	server := gin.Default()
	server.GET("/test" , func(ctx *gin.Context) {
		ctx.String(200 , "hellooo")
	})

	server.Run()
}