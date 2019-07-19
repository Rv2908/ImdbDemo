package movie

import "Imdb/model/movie"

//Movie it defines the behaviour of Movie Tables
type Movie interface {
	UpdateMovie(movie *movie.Movie) (*movie.Movie, error)
	DeleteMovie(MovieID uint) error
	AddGenre(MovieID, GenreID uint) (*movie.Movie, error)
	DeleteMovieGenre(ID uint) error
	AddMovie(movie *movie.Movie) (*movie.Movie, error)
}
