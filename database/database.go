package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

//New It will create a new database connection and returns the instance of it
func New() *sql.DB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		"imdb", "fyndimdb", "imdb")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return db
}
