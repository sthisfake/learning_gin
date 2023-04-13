package database

import "database/sql"

func GetUserFromEmail(email string) (*sql.Rows, error) {

	querry := "SELECT id , email , password  FROM users WHERE email= " + "'" + email + "'"
	result, err := db.Query(querry)

	return result, err

}