package main

import (
	"Imdb/server" 
	userController "Imdb/controllers/user"
	userRouter "Imdb/routers/user"
	"log"
	"os"

	"net/http"

	// "os"
	"strings"
	// "database/sql"
	// _ "github.com/lib/pq"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {
	logger := log.New(os.Stdout, "imdb ", log.LstdFlags|log.Lshortfile)
	// connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	// "imdb", "fyndimdb", "imdb")
	// _, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	panic(err)
	// }
	// logger.Println("Database connected")

	mux := http.NewServeMux()
	srv := server.New(mux, ":8080")
	mux.HandleFunc("/", sayHello)
	logger.Println("server starting")
	// create controllers
	uc := userController.NewUser()
	userRouter.NewUserRouter(uc).Register(mux)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
