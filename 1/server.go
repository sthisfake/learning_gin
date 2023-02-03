package main

import (
	"github.com/gin-gonic/gin"
	"sth.com/controllers"
	"sth.com/service"
)

var (
	MovieService    service.MovieService        = service.New()
	MovieController controllers.MovieController = controllers.New(MovieService)
)

func main() {

	server := gin.Default()

	server.GET("/movies", func(ctx *gin.Context) {
		ctx.JSON(200, MovieController.allMovies())
	})

	server.POST("/movies", func(ctx *gin.Context) {
		ctx.JSON(200, MovieController.saveMovie(ctx))
	})

	server.Run(":8000")
}
