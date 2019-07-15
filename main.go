package main

import (
	userController "Imdb/controllers/user"
	"Imdb/database"
	userRouter "Imdb/routers/user"
	"Imdb/server"
	"log"
	"os"

	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {
	logger := log.New(os.Stdout, "imdb ", log.LstdFlags|log.Lshortfile)
	db := database.New()
	defer db.Close()
	logger.Println("Database Created")

	mux := http.NewServeMux()
	srv := server.New(mux, ":3000")
	mux.HandleFunc("/", sayHello)
	logger.Println("server starting")
	// create controllers
	uc := userController.NewUser(db)
	userRouter.NewUserRouter(uc, logger).Register(mux)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
