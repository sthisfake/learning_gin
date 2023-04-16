package database

import (
	"database/sql"
	"time"
)

func GetUserFromUserName(userName string) (*sql.Rows, error) {

	querry := "SELECT id FROM users WHERE user_name= " + "'" + userName + "'"
	result, err := db.Query(querry)

	return result, err

}

func InserFollow(firstUserId string , secondUserId  string ) (error){
	time := time.Now().Format("2006-01-02 15:04:05")
	querry := "INSERT INTO follow (created_at , updated_at , first_user_id , followed_user_id) VALUES("+ "'" + time + "'" + ", '" + time + "'"  + ", '" + firstUserId + "'" + ", "  + "'" + secondUserId  + "'"  + ")"
	_, err := db.Query(querry)
	return err
}