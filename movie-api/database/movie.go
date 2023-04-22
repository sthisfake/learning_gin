package database

import (
	"fmt"
	"movies/models"
	"strings"
	"time"
)

func InitialMovieInsert(movie models.Movie , movies models.Movie) (error){


	/// movies to database 

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

	if err != nil {
		fmt.Printf("Error fetching data  %s\n",  err.Error())
		
	}

	// select the movie id that just inserted 

	var movieId string

	querry = "SELECT id FROM movie WHERE name= " + "'" + movie.Title + "'"
	result, err := db.Query(querry)

	if err !=nil {
		return err
	}

	for result.Next(){
		if err := result.Scan(&movieId); err != nil {
            fmt.Printf("Error fetching data  %s\n",  err.Error())
        }
	}

	// insert into famous person 
	director := movies.Director
	actor := movies.Actors
	person := director + "," + actor
	persons:= removeDuplicateStr(strings.Split(person , ",")) 


	for i:=0 ; i<len(persons) ; i++{
		querry := "INSERT INTO famous_person (full_name) VALUES("+ "'" + persons[i] + "'" + ")"
		_, err = db.Query(querry)

		if err != nil {
			fmt.Printf("Error inserting data for  famous person loop in famous persons ID %d: %s\n", i, err.Error())
			continue
		}
	}

	// select the person id that just inserted 
	// and insert into director table
	//also 
	// select the director id 
	// and insert into movie-director table

	var personId string
	var directorId string
	if strings.Contains(director , ",") {
		directors:= strings.Split(director , ",")
		for i:=0 ; i< len(directors) ; i++ {

			// director

			querry = "SELECT id FROM famous_person WHERE full_name= " + "'" + directors[i] + "'"
			result, err = db.Query(querry)
		
			for result.Next(){
				if err := result.Scan(&personId); err != nil {
					fmt.Printf("Error fetching data  %s\n",  err.Error())
				}
			}

			querry := "INSERT INTO director (person_id) VALUES("+ "'" + personId + "'" + ")"
			_, err = db.Query(querry)

			// movie director

			querry = "SELECT id FROM director WHERE person_id= " + "'" + personId + "'"
			result, err = db.Query(querry)

			if err == nil {
				for result.Next(){
					if err := result.Scan(&directorId); err != nil {
						fmt.Printf("Error fetching data  %s\n",  err.Error())
					}
				}
	
				querry = "INSERT INTO movie_director (director_id , movie_id) VALUES("+ "'" + directorId + "'" + ", " + "'" + movieId + "'" + ")"
				_, err = db.Query(querry)
			}
		
		}
	}else{

		// director

		querry = "SELECT id FROM famous_person WHERE full_name= " + "'" + director + "'"
		result, err = db.Query(querry)
	
		for result.Next(){
			if err := result.Scan(&personId); err != nil {
				fmt.Println(err)
			}
		}

		querry := "INSERT INTO director (person_id) VALUES("+ "'" + personId + "'" + ")"
		_, err = db.Query(querry)


		// movie director

		querry = "SELECT id FROM director WHERE person_id= " + "'" + personId + "'"
		result, err = db.Query(querry)

		if err == nil{
			for result.Next(){
				if err := result.Scan(&directorId); err != nil {
					fmt.Printf("Error fetching data  %s\n",  err.Error())
					}
				}
			
			querry = "INSERT INTO movie_director (director_id , movie_id) VALUES("+ "'" + directorId + "'" + ", " + "'" + movieId + "'" + ")"
			_, err = db.Query(querry)
		}
				


	}

	//insert into actor table
	var actorId string
	if strings.Contains(actor , ",") {
		actors:= strings.Split(actor , ",")
		for i:=0 ; i< len(actors) ; i++ {

			// actor
			querry = "SELECT id FROM famous_person WHERE full_name= " + "'" + actors[i] + "'"
			result, err = db.Query(querry)
		
			for result.Next(){
				if err := result.Scan(&personId); err != nil {
					fmt.Printf("Error fetching data  %s\n",  err.Error())
				}
			}

			querry := "INSERT INTO actor (person_id) VALUES("+ "'" + personId + "'" + ")"
			_, err = db.Query(querry)

			// movie actor

			querry = "SELECT id FROM actor WHERE person_id= " + "'" + personId + "'"
			result, err = db.Query(querry)

			if err == nil {
				for result.Next(){
					if err := result.Scan(&actorId); err != nil {
						fmt.Printf("Error fetching data  %s\n",  err.Error())
					}
				}
	
				querry = "INSERT INTO movie_actor (actor_id , movie_id) VALUES("+ "'" + actorId + "'" + ", " + "'" + movieId + "'" + ")"
				_, err = db.Query(querry)
			}
		
		}
	}else{

		// actor

		querry = "SELECT id FROM famous_person WHERE full_name= " + "'" + actor + "'"
		result, err = db.Query(querry)
	
		for result.Next(){
			if err := result.Scan(&personId); err != nil {
				fmt.Println(err)
			}
		}

		querry := "INSERT INTO actor (person_id) VALUES("+ "'" + personId + "'" + ")"
		_, err = db.Query(querry)


		// movie actor

		querry = "SELECT id FROM actor WHERE person_id= " + "'" + personId + "'"
		result, err = db.Query(querry)

		if err == nil{
			for result.Next(){
				if err := result.Scan(&actorId); err != nil {
					fmt.Printf("Error fetching data  %s\n",  err.Error())
					}
				}
			
			querry = "INSERT INTO movie_actor (actor_id , movie_id) VALUES("+ "'" + actorId + "'" + ", " + "'" + movieId + "'" + ")"
			_, err = db.Query(querry)
		}
				


	}
	return err

}


func removeDuplicateStr(strSlice []string) []string {
    allKeys := make(map[string]bool)
    list := []string{}
    for _, item := range strSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}
