package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), "postgres://deedsbaron:0809@localhost:5432/urlshort")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	fmt.Println("id = ", id)
	fmt.Println("Str = ", str)

	q := `INSERT INTO urls (id, longURL, shortURL) VALUES ($1, $2, $3);`
	err = conn.QueryRow(context.Background(), q, id, "https://asdasd", "http://asdasdasd").Scan()
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
}
