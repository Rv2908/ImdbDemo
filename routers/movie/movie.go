package movie

import (
	interfaces "Imdb/interfaces/movie"
	movie "Imdb/model/movie"
	middleware "Imdb/token"
	"encoding/json"
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

// MethodHandlerMovie Identifies the Method and call the api accordingly
func (mv Movie) MethodHandlerMovie(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		mv.getMovie(w, r)

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

// MethodHandlerGenre Identifies the Method and call the api accordingly
func (mv Movie) MethodHandlerGenre(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		mv.addMovieGenre(w, r)
	case http.MethodDelete:
		mv.deleteMovieGenre(w, r)

	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Unknown Method Type")
	}
}

//Register This method consist of all the routes for the movie
func (mv Movie) Register(s *http.ServeMux) {
	s.Handle("/movie/", mv.Logger(middleware.JWTMiddleware(middleware.AdminMiddleware(http.HandlerFunc(mv.MethodHandlerMovie)))))
	s.Handle("/genre/", mv.Logger(middleware.JWTMiddleware(middleware.AdminMiddleware(http.HandlerFunc(mv.MethodHandlerGenre)))))
	s.HandleFunc("/movie", mv.Logger(middleware.JWTMiddleware(mv.getMovies)))
}

func (mv Movie) getMovies(w http.ResponseWriter, r *http.Request) {

	pageNo := getQueryUint(r, "pageNo")
	pageSize := getQueryUint(r, "pageSize")
	searchBy := getQueryString(r, "search")

	movie, err := mv.movieController.GetAllMovie(pageNo, pageSize, searchBy)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	b, _ := json.Marshal(movie)

	w.Write(b)
}

func (mv Movie) getMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called getMovie")
	MovieID := getQueryKey(r)
	if MovieID == 0 {
		w.Write([]byte("Movie ID is missing"))
	}

	movie, err := mv.movieController.GetMovie(MovieID)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	b, _ := json.Marshal(movie)

	w.Write(b)

}

func (mv Movie) deleteMovieGenre(w http.ResponseWriter, r *http.Request) {
	ID := getQueryKey(r)
	fmt.Println("Delete Movie_ Genre Having ID ", ID)
	if err := mv.movieController.DeleteMovieGenre(ID); err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Movie Genre Deleted Successfully"))
}

func (mv Movie) addMovieGenre(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		MovieID uint   `json:"movie_id"`
		Genres  []uint `json:"genres"`
	}

	request := &Request{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(request)

	mve, err := mv.movieController.AddGenre(request.MovieID, request.Genres)
	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	b, _ := json.Marshal(mve)
	w.Write(b)

}

func (mv Movie) addMovie(w http.ResponseWriter, r *http.Request) {

	movie := &movie.MovieRequest{}
	if err := json.NewDecoder(r.Body).Decode(movie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(movie)
	mve, err := mv.movieController.AddMovie(movie)
	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	b, _ := json.Marshal(mve)
	w.Write(b)
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

func (mv Movie) updateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Called updateMovie")
	movie := &movie.MovieUpdate{}
	if err := json.NewDecoder(r.Body).Decode(movie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mve, err := mv.movieController.UpdateMovie(movie)
	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	b, _ := json.Marshal(mve)
	w.Write(b)
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

func getQueryString(r *http.Request, param string) string {
	keys, ok := r.URL.Query()[param]
	if !ok || len(keys[0]) < 1 {
		return ""
	}

	return keys[0]
}

func getQueryUint(r *http.Request, param string) uint {
	keys, ok := r.URL.Query()[param]
	if !ok || len(keys[0]) < 1 {
		return 0
	}

	ID, err := strconv.ParseUint(keys[0], 10, 32)
	if err != nil {
		return 0
	}

	return uint(ID)
}

func getQueryfloat(r *http.Request, param string) float32 {
	keys, ok := r.URL.Query()[param]
	if !ok || len(keys[0]) < 1 {
		return 0
	}

	ID, err := strconv.ParseFloat(keys[0], 10)
	if err != nil {
		return 0
	}

	return float32(ID)
}
