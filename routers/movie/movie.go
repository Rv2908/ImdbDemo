package movie

import (
	interfaces "Imdb/interfaces/movie"
	"Imdb/model/movie"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Movie contains the instance of movie controller and logger instance
type Movie struct {
	movieController interfaces.Movie
	logger          *log.Logger
}

//Logger log all the incoming method
func (mv Movie) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer mv.logger.Printf("New Request in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

//NewMovieRouter Return instance of Movie Router
func NewMovieRouter(movieController interfaces.Movie, logger *log.Logger) Movie {
	return Movie{
		movieController: movieController,
		logger:          logger,
	}
}

//Register This method consist of all the routes for the movie
func (mv Movie) Register(s *http.ServeMux) {
	s.HandleFunc("/movie", mv.Logger(mv.addMovie))
	s.HandleFunc("/movie/", mv.Logger(mv.deleteMovie))
}

func (mv Movie) addMovie(w http.ResponseWriter, r *http.Request) {
	movie := &movie.Movie{}
	if err := json.NewDecoder(r.Body).Decode(movie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(movie)
	if _, err := mv.movieController.AddMovie(movie); err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Movie Created Successfully"))

}

func (mv Movie) deleteMovie(w http.ResponseWriter, r *http.Request) {
	movie := &movie.Movie{}
	if err := json.NewDecoder(r.Body).Decode(movie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(movie)
	if err := mv.movieController.DeleteMovie(1); err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Movie Deleted Successfully"))

}
