package database

import (
	"movies/models"
	"time"
)

func InitialInsert(movie models.Movie) (error){

	time := time.Now().Format("2006-01-02 15:04:05")

	var genre string

	genre = movie.Genres[0]

	for	i:=1 ; i< len(movie.Genres)  ; i++{
		genre = genre + " , "  + movie.Genres[i]
	}

	querry := "INSERT INTO movie (name ,year , plot , country , poster , created_at , updated_at , genres , runtime , imdb_id) VALUES("+ 
	"'" + movie.Title + "'" + ", '" + movie.Year + "'"  + ", '" + movie.Plot + "'" + ", "  + "'" + movie.Country  + "'"  + ", "  + "'" +
	 movie.Poster  + "'"  +", "  + "'" + time  + "'"  +", "  + "'" + time  + "'"  +", "  + "'" + genre  + "'"  +", "  + "'" +
	  movie.Runtime  + "'"  +", "  + "'" + movie.IMDBID  + "'"  + ")" 	
	_, err := db.Query(querry)
  
	return err

}