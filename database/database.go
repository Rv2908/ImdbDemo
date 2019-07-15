package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//New It will create a new database connection and returns the instance of it
func New() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"imdb.csnbws2ql3w6.us-east-2.rds.amazonaws.com", 5432, "imdb", "fyndimdb", "imdb")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}
