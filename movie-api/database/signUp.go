package database

import (
	"movies/models"
	"time"
)

func CreateUser( user  models.User ) ( error) {

  time := time.Now().Format("2006-01-02 15:04:05")

  querry := "INSERT INTO users (created_at , updated_at ,  email , password) VALUES("+ "'" + time + "'" + ", '" + time + "'"  + ", '" + user.Email + "'" + ", "  + "'" + user.Password  + "'"  + ")" 	
  _, err := db.Query(querry)

  return err


}