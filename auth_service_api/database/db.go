package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var Pool *pgxpool.Pool

// InitDB initializes a connection pool to a Postgres database.
// It requires a dataSourceName string which contains the connection parameters.
// The function will log.Fatal if it fails to establish a connection or ping the database.
func InitDB(dataSourceName string) {
	var err error
	Pool, err = pgxpool.Connect(context.Background(), dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	err = Pool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to the database!")
}
