package movie

import "Imdb/model/movie"

//Movie it defines the behaviour of Movie Tables
type Movie interface {
	UpdateMovie(movie *movie.MovieUpdate) (movie.MovieGet, error)
	AddGenre(MovieID uint, GenreID []uint) (movie.MovieGet, error)
	AddMovie(movie *movie.MovieRequest) (movie.MovieGet, error)
	GetMovie(MovieID uint) (movie.MovieGet, error)
	DeleteMovie(MovieID uint) error
	DeleteMovieGenre(ID uint) error
	GetAllMovie(PageNo, PageSize uint, SearchBy string) ([]movie.MovieGet, error)
}
