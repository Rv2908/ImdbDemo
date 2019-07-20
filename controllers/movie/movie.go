package movie

import (
	movie "Imdb/model/movie"
	"database/sql"
	"time"
)

//Movie it contains the instance of database
type Movie struct {
	db *sql.DB
}

//NewMovie this will return the instance of movie with database instance in it
func NewMovie(db *sql.DB) Movie {
	return Movie{db}
}

// UpdateMovie This function will update the movie
func (mv Movie) UpdateMovie(movie *movie.Movie) (*movie.Movie, error) {
	sqlStatement := `UPDATE movies
					SET name = $2, 
					updated_at = $3, 
					popularity = $4, 
					imdb_score=$5, 
					director = $6
					WHERE id = $1;`
	_, err := mv.db.Exec(sqlStatement, movie.ID, movie.Name, time.Now(), movie.Popularity, movie.ImdbScore, movie.Director)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

// DeleteMovie it deletes the movie and refernce of movie in mapping table
func (mv Movie) DeleteMovie(MovieID uint) error {
	tx, err := mv.db.Begin()
	if err != nil {
		return err
	}

	sqlStatementDeleteMovie := `
					DELETE FROM movies
					WHERE id = $1;`
	_, err = tx.Exec(sqlStatementDeleteMovie, MovieID)
	if err != nil {
		tx.Rollback()
		return err
	}

	sqlStatementDeleteMapping := `DELETE FROM movie_genres
	WHERE movie_id = $1`
	_, err = tx.Exec(sqlStatementDeleteMapping, MovieID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// AddGenre add new genre to the movie
func (mv Movie) AddGenre(MovieID, GenreID uint) (*movie.Movie, error) {
	genreMovie := new(movie.Movie)

	var id uint
	sqlStatementFindRecord := "select id from movie_genres where movie_id=$1 and genre_id= $2"
	row := mv.db.QueryRow(sqlStatementFindRecord, MovieID, GenreID)
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		sqlStatementAddGenre := `INSERT INTO movie_genres(genre_id,movie_id) VALUES ($1, $2)`
		_, err := mv.db.Exec(sqlStatementAddGenre, GenreID, MovieID)
		if err != nil {
			return nil, err
		}

	default:
		return nil, err
	}

	return genreMovie, nil
}

//DeleteMovieGenre Delete the movie genre
func (mv Movie) DeleteMovieGenre(ID uint) error {
	sqlStatement := `
					DELETE FROM movie_genres
					WHERE id = $1;`
	_, err := mv.db.Exec(sqlStatement, ID)
	if err != nil {
		return err
	}
	return nil
}

// AddMovie this method adds the new movie
func (mv Movie) AddMovie(mve *movie.Movie) (*movie.Movie, error) {
	return new(movie.Movie), nil
}
