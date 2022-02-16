CREATE TABLE urls (
      id bigserial PRIMARY KEY,
      longURL VARCHAR ( 50 ) UNIQUE NOT NULL,
      shortURL VARCHAR ( 50 ) NOT NULL
);