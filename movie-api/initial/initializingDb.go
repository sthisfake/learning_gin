package initial

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"movies/database"
	"movies/models"
	"net/http"
	"os"
)

func FillingFamousPersonTable(){
	database.StartDb()

	for i := 124; i < 250; i++ {

		
		jsonData, err := ioutil.ReadFile("movies.json")
		if err != nil {
			fmt.Println("Error reading movies data from file:", err)
			os.Exit(1)
		}
	
		movies := []models.Movie{}
		err = json.Unmarshal(jsonData, &movies)
		if err != nil {
			fmt.Println("Error decoding movies data from JSON:", err)
			os.Exit(1)
		}

		newMovie := models.Movie{
			Director: trimQuotes(movies[i].Director),
			Actors: trimQuotes(movies[i].Actors),
		}

		// insert into movie table

		err = database.InitialFamousPersonsInsert(newMovie)

		if(err != nil){
			fmt.Printf("movie id %d  NOT done \n" , i)
			fmt.Println("*******************************")
			fmt.Println(err)
			fmt.Println("*******************************")
		} else{
			fmt.Printf("movie id %d done \n" , i)
		}

	}

	database.CloseDb()
}

func FillingTheMovieTable() {

	database.StartDb()

	for i := 0; i < 1; i++ {

		// all movies in movies.json file into a slice of array

		jsonData, err := ioutil.ReadFile("movies.json")
		if err != nil {
			fmt.Println("Error reading movies data from file:", err)
			os.Exit(1)
		}
	
		// decode JSON data into slice of Movie structs
		movies := []models.Movie{}
		err = json.Unmarshal(jsonData, &movies)
		if err != nil {
			fmt.Println("Error decoding movies data from JSON:", err)
			os.Exit(1)
		}

		newMovie := models.Movie{
			Title: trimQuotes(movies[i].Title),
			Poster: trimQuotes(movies[i].Poster) ,
			Year: trimQuotes(movies[i].Year),
			Runtime: trimQuotes(movies[i].Runtime),
			Plot: trimQuotes(movies[i].Plot),
			Country: trimQuotes(movies[i].Country),
			IMDBID: trimQuotes(movies[i].IMDBID),
			Genres: movies[i].Genres,
		}

		// insert into movie table

		err = database.InitialMovieInsert(newMovie , movies[i])

		if(err != nil){
			fmt.Printf("movie id %d  NOT done \n" , i)
			fmt.Println("*******************************")
			fmt.Println(err)
			fmt.Println("*******************************")
		} else{
			fmt.Printf("movie id %d done \n" , i)
		}

	}

	database.CloseDb()
}


func GettingMoviesFromApi() {
	baseURL := "https://moviesapi.ir/api/v1/movies/"
	movies := []models.Movie{}

	for i := 1; i <= 250; i++ {
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

		movie := models.Movie{}
		err = json.Unmarshal(body, &movie)
		if err != nil {
			fmt.Printf("Error decoding response body for ID %d: %s\n", i, err.Error())
			continue
		}

		movies = append(movies, movie)
	}

	jsonData, err := json.Marshal(movies)
	if err != nil {
		fmt.Println("Error encoding movies data to JSON:", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile("movies.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing movies data to file:", err)
		os.Exit(1)
	}

	fmt.Println("Data saved to movies.json")
}

func trimQuotes(s string) string {

	for i := 0; i < len(s); i++ {
		if s[i] == '\'' {
			s = strRemoveAt(s, i, 1)
		}
	}

	return s
}

func strRemoveAt(s string, index, length int) string {
	return s[:index] + s[index+length:]
}