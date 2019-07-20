package movie

import (
	interfaces "Imdb/interfaces/movie"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

// MethodHandler Identifies the Method and call the api accordingly
func (mv Movie) MethodHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		ID := getQueryKey(r)
		if ID == 0 {
			mv.getMovies(w, r)
		} else {
			mv.getMovie(w, r)
		}

	case http.MethodPost:
		mv.addMovie(w, r)

	case http.MethodPut:
		mv.updateMovie(w, r)

	case http.MethodDelete:
		mv.deleteMovie(w, r)

	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Unknown Method Type")
	}
}

//Register This method consist of all the routes for the movie
func (mv Movie) Register(s *http.ServeMux) {
	s.Handle("/movie/", mv.Logger(http.HandlerFunc(mv.MethodHandler)))
}

func (mv Movie) addMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called addMovie")
	w.Write([]byte("Success"))
	// movie := &movie.Movie{}
	// if err := json.NewDecoder(r.Body).Decode(movie); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// fmt.Println(movie)
	// if _, err := mv.movieController.AddMovie(movie); err != nil {
	// 	fmt.Println(err.Error())
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	// w.Write([]byte("Movie Created Successfully"))
}

func (mv Movie) deleteMovie(w http.ResponseWriter, r *http.Request) {
	movieID := getQueryKey(r)
	fmt.Println("Delete Movie Having ID ", movieID)
	if err := mv.movieController.DeleteMovie(movieID); err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Movie Deleted Successfully"))
}

func (mv Movie) getMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called getMovies")
	w.Write([]byte("Success"))

}

func (mv Movie) updateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called updateMovie")
	w.Write([]byte("Success"))
}

func (mv Movie) getMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called getMovie")
	w.Write([]byte("Success"))
	// MovieID := getQueryKey(r)
	// if MovieID == 0 {
	// 	w.Write([]byte("Movie ID is missing"))
	// }

}

func getQueryKey(r *http.Request) uint {
	keys, ok := r.URL.Query()["ID"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Movie ID is missing")
		return 0
	}

	movieID, err := strconv.ParseUint(keys[0], 10, 32)
	if err != nil {
		return 0
	}

	return uint(movieID)
}
