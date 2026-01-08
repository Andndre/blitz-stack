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
	CREATE TABLE IF NOT EXISTS dealers (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		city TEXT NOT NULL
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func Seed(db *sql.DB) {
	// Seed data if empty logic
	var count int
	db.QueryRow("SELECT COUNT(*) FROM dealers").Scan(&count)
	
	if count == 0 {
		query := `INSERT INTO dealers (name, city) VALUES 
			('Kantor Pusat', 'Jalan Kartini Nomor 71 hingga 75, Kelurahan Dangin Puri Kaja, Denpasar Utara'),
			('Kantor Cabang WR Supratman', 'Jalan WR Supratman Nomor 130, Denpasar'),
			('Kantor Cabang Tuban', 'Jalan Raya Tuban Nomor 100X, Tuban, Kabupaten Badung (wilayah Kuta Selatan)'),
			('Pos Penjualan Pemogan', 'Berlokasi di daerah Pemogan')`
		
		_, err := db.Exec(query)
		if err != nil {
			log.Printf("Error seeding data: %v\n", err)
		} else {
			fmt.Println("Seeded database with initial dealer data.")
		}
	} else {
		fmt.Println("Database already contains data. Skipping seed.")
	}
}