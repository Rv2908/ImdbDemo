CREATE TABLE users
(
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  first_name TEXT,
  last_name TEXT,
  email TEXT UNIQUE NOT NULL,
  password TEXT,
  role_id INT
);

CREATE TABLE roles
(
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  role TEXT UNIQUE NOT NULL
);

CREATE TABLE genres
(
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  genre TEXT UNIQUE NOT NULL
);


CREATE TABLE movies
(
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  popularity NUMERIC (4, 2) NOT NULL,
  director TEXT UNIQUE NOT NULL,
  name TEXT NOT NULL,
  imdb_score NUMERIC (3, 1) NOT NULL,
);


CREATE TABLE movie_genres
(
  id SERIAL PRIMARY KEY,
  movie_id INT REFERENCES movies (id) NOT NULL,
  genre_id INT REFERENCES genre (id) NOT NULL,
);