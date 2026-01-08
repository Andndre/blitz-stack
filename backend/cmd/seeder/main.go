package main

import (
	"log"

	"dealer-heronusa/backend/internal/database"
)

func main() {
	// Initialize Database Connection
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Could not connect to database for seeding: %v", err)
	}
	defer db.Close()

	log.Println("Starting data seeding...")
	database.Seed(db)
	log.Println("Seeding process finished.")
}
