package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	//load html glob

	server.LoadHTMLGlob("./templates/*")

	//load html files

	// server.LoadHTMLFiles(
	// 	"./templates/home.html",
	// 	"./templates/about.html",
	// )

	server.GET("/home", func(ctx *gin.Context) {
		ctx.HTML(200, "home.html", gin.H{
			"title":     "testing",
			"something": "bla blab lbaba",
		})
	})

	server.GET("/about", func(ctx *gin.Context) {
		ctx.HTML(200, "about.html", nil)
	})

	// redirect

	server.GET("/redirect", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusTemporaryRedirect, "/home")
	})

	server.Run(":8080")

}
