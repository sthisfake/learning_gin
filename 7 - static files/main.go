package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	sth := gin.Default()

	// sth.static

	sth.Static("/test", "./assets")

	//sth.staticFs

	sth.StaticFS("/sth", http.Dir("./assets"))

	//sth.staticFile

	sth.StaticFile("/how", "./assets/something.txt")

	sth.Run(":8080")

}
