package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func InitTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS items (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func Seed(db *sql.DB) {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM items").Scan(&count)
	
	if count == 0 {
		query := `INSERT INTO items (title, description) VALUES 
			('First Blitz Item', 'This is a sample item from the Blitz Stack seeder.க்கான'),
			('Scalable Backend', 'Go provides amazing performance for your API.'),
			('Fast Frontend', 'Next.js and Bun ensure a snappy user experience.')`
		
		_, err := db.Exec(query)
		if err != nil {
			log.Printf("Error seeding data: %v\n", err)
		} else {
			fmt.Println("Seeded database with initial Blitz data.")
		}
	} else {
		fmt.Println("Database already contains data. Skipping seed.")
	}
}
