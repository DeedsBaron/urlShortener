CREATE DATABASE urlshort;
\c urlshort;
CREATE TABLE urls (
    id NUMERIC PRIMARY KEY,
    longurl TEXT UNIQUE NOT NULL,
    shorturl TEXT NOT NULL
);
