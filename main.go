package main

import (
	movieController "Imdb/controllers/movie"
	userController "Imdb/controllers/user"
	"Imdb/database"
	movieRouter "Imdb/routers/movie"
	userRouter "Imdb/routers/user"
	"Imdb/server"
	token "Imdb/token"
	"log"
	"os"

	"net/http"
)

func main() {
	logger := log.New(os.Stdout, "imdb ", log.LstdFlags|log.Lshortfile)

	logger.Println("CONNECT TO DATABASE")
	db, err := database.New()
	if err != nil {
		log.Panic("COULDN'T CONNECT TO DATABASE " + err.Error())
	}
	defer db.Close()

	logger.Println("START UP SERVER")

	mux := http.NewServeMux()
	srv := server.New(mux, ":3000")

	// generate and set random token key
	token.SetAccessTokenSecret()

	// create controllers
	uc := userController.NewUser(db)
	mv := movieController.NewMovie(db)

	// define routers
	userRouter.NewUserRouter(uc, logger).Register(mux)
	movieRouter.NewMovieRouter(mv, logger).Register(mux)

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
