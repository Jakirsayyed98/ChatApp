package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectPostgres() (*sql.DB, error) {
	postgresURL := os.Getenv("PG_CONNECION_STRING")
	fmt.Printf("Connecting to Postgres at %s\n", postgresURL)
	sql, err := sql.Open("postgres", postgresURL)
	if err != nil {
		return nil, err
	}

	if err := sql.Ping(); err != nil {
		return nil, err
	}

	DB = sql

	return sql, nil
}

func GetDB() *sql.DB {
	return DB
}
