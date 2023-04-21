package initial

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"movies/database"
	"movies/models"
	"net/http"
)

func FillingTheMovieTable() {
	baseURL := "https://moviesapi.ir/api/v1/movies/"

	database.StartDb()
	
	for i := 206; i <= 250; i++ {

		url := fmt.Sprintf("%s%d", baseURL, i)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error fetching data for ID %d: %s\n", i, err.Error())
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body for ID %d: %s\n", i, err.Error())
			continue
		}


		var movie models.Movie
		err = json.Unmarshal(body, &movie)
		if err != nil {
			log.Fatal(err)
		}

		

		newMovie := models.Movie{
			Title: movie.Title,
			Poster: movie.Poster ,
			Year: movie.Year,
			Runtime: movie.Runtime,
			Plot: movie.Plot,
			Country: movie.Country,
			IMDBID: movie.IMDBID,
			Genres: movie.Genres,
		}

		err = database.InitialInsert(newMovie)

		if err != nil {
			fmt.Printf("Error fetching data for ID %d: %s\n", i, err.Error())
			continue
		}


		fmt.Printf("movie id %d done \n" , i)


		
	}

	database.CloseDb()
}
