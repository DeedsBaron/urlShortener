CREATE TABLE urls (
<<<<<<< HEAD
      id bigserial PRIMARY KEY,
=======
      id NUMERIC PRIMARY KEY,
>>>>>>> ed8f4a1 (postgresql container is configured and working)
      longURL VARCHAR ( 50 ) UNIQUE NOT NULL,
      shortURL VARCHAR ( 50 ) NOT NULL
);