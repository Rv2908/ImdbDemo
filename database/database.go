package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//New It will create a new database connection and returns the instance of it
func New() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"imdb.csnbws2ql3w6.us-east-2.rds.amazonaws.com", 5432, "imdb", "fyndimdb", "imdb")
	return sql.Open("postgres", psqlInfo)

}
