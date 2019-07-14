package main

import (
	"Imdb/server"
	"net/http"
	"strings"
	"os"
	"log"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {

	logger := log.New(os.Stdout, "imdb ", log.LstdFlags|log.Lshortfile)
	mux := http.NewServeMux()

	srv := server.New(mux, ":8080")
	mux.HandleFunc("/", sayHello)
	logger.Println("server starting")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
