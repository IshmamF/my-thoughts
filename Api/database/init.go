package database 

import (
	"fmt"
	"context"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type DB struct {
	conn *pgxpool.Pool
}


func (d *DB) Connect () {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	dbUrl := os.Getenv("DB_URL")

	conn, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	d.conn = conn
}

func (d *DB) Ping() error {
	return d.conn.Ping(context.Background())
}
