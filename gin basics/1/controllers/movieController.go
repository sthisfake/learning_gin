package controllers

import (
	"github.com/gin-gonic/gin"
	"sth.com/models"
	"sth.com/service"
)

type MovieController interface {
	allMovies() []models.Movie
	saveMovie(ctx *gin.Context)
}

type controller struct {
	servicess service.MovieService
}

func New(service service.MovieService) MovieController {
	return &controller{
		servicess: service,
	}
}

func (c *controller) allMovies() []models.Movie {
	return c.servicess.allMovies()
}

func (c *controller) saveMovie(ctx *gin.Context) {
	var movie models.Movie
	ctx.BindJSON((&movie))
	c.servicess.saveMovie(movie)
}
