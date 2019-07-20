package movie

import (
	"encoding/json"
	"time"
)

// Movie contains information about the movie
type Movie struct {
	ID         uint       `json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	Popularity float32    `json:"popularity"`
	Director   string     `json:"director"`
	Name       string     `json:"name"`
	ImdbScore  float32    `json:"imdb_score"`
	Genre      []Genre    `json:"genre"`
}

//MovieRequest Add Request
type MovieRequest struct {
	Popularity float32 `json:"popularity"`
	Director   string  `json:"director"`
	Name       string  `json:"name"`
	ImdbScore  float32 `json:"imdb_score"`
	Genre      []uint  `json:"genres"`
}

type MovieGet struct {
	ID         uint            `json:"id"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	DeletedAt  *time.Time      `json:"deleted_at"`
	Popularity float32         `json:"popularity"`
	Director   string          `json:"director"`
	Name       string          `json:"name"`
	ImdbScore  float32         `json:"imdb_score"`
	Genre      json.RawMessage `json:"genre"`
}

type MovieUpdate struct {
	ID         uint    `json:"id"`
	Popularity float32 `json:"popularity"`
	Director   string  `json:"director"`
	Name       string  `json:"name"`
	ImdbScore  float32 `json:"imdb_score"`
}

// Genre contains information about genre of the movie
type Genre struct {
	ID    uint   `json:"id"`
	Genre string `json:"genre"`
}

// MovieGenre mapping table of movie and genre
type MovieGenre struct {
	ID      uint `json:"id"`
	MovieID uint `json:"movie_id"`
	GenreID uint `json:"genre_id"`
}
