package main

import (
	"log"

	"blitz-stack/backend/internal/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Could not connect to database for seeding: %v", err)
	}
	defer db.Close()

	log.Println("Starting Blitz data seeding...")
	database.Seed(db)
	log.Println("Seeding process finished.")
}