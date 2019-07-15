package model

// Movie contains information about the movie
type Movie struct {
	ID         uint
	Name       string
	Director   string
	Popularity float32
	Rating     float32
	Genre      []Genre
}

// Genre contains information about genre of the movie
type Genre struct {
	ID    uint
	Genre string
}
