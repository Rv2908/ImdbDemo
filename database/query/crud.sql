-- Get Details of single movie
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
    ) t) 
from movies
where movies.id = 1

-- _____________________________________________________________________________________________

-- Get all movies lazy loading

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
    ) t) 
from movies
offset 0 limit 10


-- _____________________________________________________________________________________________

