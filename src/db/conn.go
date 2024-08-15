package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")

	DATABASE_URL := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_HOST, POSTGRES_DB)

	if DATABASE_URL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	fmt.Print(DATABASE_URL)
	db, err := sql.Open("postgres", DATABASE_URL)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db, nil
}
