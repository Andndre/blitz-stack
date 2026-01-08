package main

import (
	"fmt"
	"log"
	"net/http"

	"blitz-stack/backend/internal/database"
	"blitz-stack/backend/internal/handlers"
	"blitz-stack/backend/internal/middleware"
)

func main() {
	// Initialize Database
	db, err := database.Connect()
	if err != nil {
		log.Printf("Warning: Could not connect to database: %v\n", err)
	} else {
		fmt.Println("Successfully connected to the database!")
		defer db.Close()
		
		// Initialize Tables
		database.InitTable(db)
	}

	// Create a new ServeMux (Router)
	mux := http.NewServeMux()

	// Setup Routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Blitz Stack API!")
	})

	mux.HandleFunc("/health", handlers.HealthCheck(db))
	mux.HandleFunc("/api/items", handlers.GetItems(db))

	// Apply Global Middleware (Rate Limiting)
	wrappedMux := middleware.LimitMiddleware(mux)

	// Start Server
	fmt.Println("Blitz Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", wrappedMux); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}