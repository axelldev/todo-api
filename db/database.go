package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// NewDatabase creates a new database connection and returns it
func NewDatabase() *sql.DB {
	url := os.Getenv("DB_URL")
	log.Println(url)
	if url == "" {
		log.Fatal("DB_URL is not set")
	}
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		fmt.Println()
		log.Fatal(err)
	}
	return db
}
