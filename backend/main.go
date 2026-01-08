package main

import (
	"fmt"
	"log"
	"net/http"

	"dealer-heronusa/backend/internal/database"
	"dealer-heronusa/backend/internal/handlers"
	"dealer-heronusa/backend/internal/middleware"
)

func main() {
	// Initialize Database
	db, err := database.Connect()
	if err != nil {
		log.Printf("Warning: Could not connect to database: %v\n", err)
	} else {
		fmt.Println("Successfully connected to the database!")
		defer db.Close()
		
		// Initialize Tables & Seed Data
		database.InitTable(db)
	}

	// Create a new ServeMux (Router)
	mux := http.NewServeMux()

	// Setup Routes on the mux
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go Backend!")
	})

	mux.HandleFunc("/health", handlers.HealthCheck(db))
	mux.HandleFunc("/api/dealers", handlers.GetDealers(db))

	// Apply Global Middleware (Rate Limiting) to all routes
	wrappedMux := middleware.LimitMiddleware(mux)

	// Start Server with the wrapped handler
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", wrappedMux); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
