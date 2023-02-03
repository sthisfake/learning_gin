package service

import "sth.com/models"

type MovieService interface {
	saveMovie(models.Movie)
	allMovies() []models.Movie
}

type movieService struct {
	movies []models.Movie
}

func New() MovieService {

	return &movieService{
		movies: []models.Movie{},
	}
}

func (service *movieService) saveMovie(movie models.Movie) {

	service.movies = append(service.movies, movie)

}

func (service *movieService) allMovies() []models.Movie {
	return service.movies

}
