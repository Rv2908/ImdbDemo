package movie

import (
	movie "Imdb/model/movie"
	"database/sql"
	"fmt"
	"strconv"
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
func (mv Movie) UpdateMovie(movieUpdate *movie.MovieUpdate) (movie.MovieGet, error) {

	var movie movie.MovieGet

	sqlStatement := `UPDATE movies
					SET name = $2, 
					updated_at = $3, 
					popularity = $4, 
					imdb_score=$5, 
					director = $6
					WHERE id = $1;`
	_, err := mv.db.Exec(sqlStatement, movieUpdate.ID, movieUpdate.Name, time.Now(), movieUpdate.Popularity, movieUpdate.ImdbScore, movieUpdate.Director)
	if err != nil {
		return movie, err
	}

	movie, err = mv.GetMovie(movieUpdate.ID)
	if err != nil {
		return movie, err
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
func (mv Movie) AddGenre(MovieID uint, GenreID []uint) (movie.MovieGet, error) {
	var movie movie.MovieGet

	tx, err := mv.db.Begin()
	if err != nil {
		return movie, err
	}

	for _, genreID := range GenreID {
		if err := addGenreToMovie(tx, uint(MovieID), genreID); err != nil {
			tx.Rollback()
			return movie, err
		}
	}

	tx.Commit()

	movie, err = mv.GetMovie(MovieID)
	if err != nil {
		return movie, err
	}

	return movie, nil
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
func (mv Movie) AddMovie(mve *movie.MovieRequest) (movie.MovieGet, error) {
	var movie movie.MovieGet

	tx, err := mv.db.Begin()
	if err != nil {
		return movie, err
	}

	sqlStatement := `INSERT INTO movies (created_at, updated_at,popularity, director, name, imdb_score) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	MovieID := 0
	errMovie := tx.QueryRow(sqlStatement, time.Now(), time.Now(), mve.Popularity, mve.Director, mve.Name, mve.ImdbScore).Scan(&MovieID)
	if errMovie != nil {
		tx.Rollback()
		return movie, errMovie
	}

	for _, genreID := range mve.Genre {
		if err := addGenreToMovie(tx, uint(MovieID), genreID); err != nil {
			tx.Rollback()
			return movie, err
		}
	}

	tx.Commit()

	movie, err = mv.GetMovie(uint(MovieID))
	if err != nil {
		return movie, err
	}

	return movie, nil
}

func addGenreToMovie(tx *sql.Tx, MovieID, GenreID uint) error {

	var id uint
	sqlStatementFindRecord := "select id from movie_genres where movie_id=$1 and genre_id= $2"
	row := tx.QueryRow(sqlStatementFindRecord, MovieID, GenreID)
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		sqlStatementAddGenre := `INSERT INTO movie_genres(genre_id,movie_id) VALUES ($1, $2)`
		_, err := tx.Exec(sqlStatementAddGenre, GenreID, MovieID)
		if err != nil {
			return err
		}
	case nil:
		fmt.Println("Genre Already Present")

	default:
		tx.Rollback()
		return err
	}

	return nil
}

// GetMovie Retruns the Movie of the specified ID
func (mv Movie) GetMovie(MovieID uint) (movie.MovieGet, error) {

	var singleMovie movie.MovieGet

	sqlStatement := `
	select 
	movies.id, 
	movies.name, 
	movies.director,
	movies.popularity, 
	movies.imdb_score,
	(select array_to_json(array_agg(row_to_json(t)))
		from (
		 select genres.id, 
				 genres.genre 
		 from  movie_genres 
		 left join genres on genres.id = movie_genres.genre_id 
		 where movie_genres.movie_id =movies.id
		) t) as genre
	from movies
	where movies.id = $1`
	if err := mv.db.QueryRow(sqlStatement, MovieID).Scan(&singleMovie.ID, &singleMovie.Name, &singleMovie.Director, &singleMovie.Popularity, &singleMovie.ImdbScore, &singleMovie.Genre); err != nil {
		return singleMovie, err
	}

	return singleMovie, nil
}

// GetAllMovie Get the movie list by lazy loading
func (mv Movie) GetAllMovie(PageNo, PageSize uint, SearchBy string) ([]movie.MovieGet, error) {

	var movieList []movie.MovieGet
	sqlStatement := `select 
	movies.id, 
	movies.name, 
	movies.director,
	movies.popularity, 
	movies.imdb_score,
	(select array_to_json(array_agg(row_to_json(t)))
		from (
		 select genres.id, 
				 genres.genre 
		 from  movie_genres 
		 left join genres on genres.id = movie_genres.genre_id 
		 where movie_genres.movie_id =movies.id
		) t) as genre
	from movies`

	clause := " WHERE "
	searchCondition := ""

	if SearchBy != "" {
		search := "'%" + SearchBy + "%'"
		searchCondition = clause +
			`movies.name ilike ` + search + ` OR ` +
			`movies.director ilike ` + search
		sqlStatement += searchCondition
	}

	if PageNo != 0 && PageSize != 0 {
		PS := strconv.FormatUint(uint64(PageSize), 10)
		OFFSET := strconv.FormatUint(uint64((PageNo-1)*PageSize), 10)
		sqlStatement += ` limit ` + PS + ` offset ` + OFFSET
	}

	fmt.Println(sqlStatement)

	rows, err := mv.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var singleMovie movie.MovieGet
		if err := rows.Scan(&singleMovie.ID, &singleMovie.Name, &singleMovie.Director, &singleMovie.Popularity, &singleMovie.ImdbScore, &singleMovie.Genre); err != nil {
			return nil, err
		}
		movieList = append(movieList, singleMovie)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return movieList, nil

}
