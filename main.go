package main

import (
	userController "Imdb/controllers/user"
	userRouter "Imdb/routers/user"
	"Imdb/database"
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
	srv := server.New(mux, ":8080")
	mux.HandleFunc("/", sayHello)
	logger.Println("server starting")
	// create controllers
	uc := userController.NewUser(db)
	userRouter.NewUserRouter(uc).Register(mux)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
