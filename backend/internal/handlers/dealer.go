package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"fmt"

	"dealer-heronusa/backend/models"
)

func HealthCheck(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Database Connection: OK")
	}
}

func GetDealers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, city FROM dealers")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		dealers := []models.Dealer{}
		
		for rows.Next() {
			var d models.Dealer
			if err := rows.Scan(&d.ID, &d.Name, &d.City); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			dealers = append(dealers, d)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dealers)
	}
}
