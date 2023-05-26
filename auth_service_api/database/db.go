package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var Pool *pgxpool.Pool

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

	fmt.Println("Successfully connected to the database!")
}
