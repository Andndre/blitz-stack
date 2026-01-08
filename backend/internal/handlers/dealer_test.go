package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"dealer-heronusa/backend/models"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetDealers(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Mock data
	rows := sqlmock.NewRows([]string{"id", "name", "city"}).
		AddRow(1, "Dealer Test A", "City A").
		AddRow(2, "Dealer Test B", "City B")

	// Expect the query to be executed
	mock.ExpectQuery("SELECT id, name, city FROM dealers").WillReturnRows(rows)

	// Create a request
	req, err := http.NewRequest("GET", "/api/dealers", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := GetDealers(db)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Assertions
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var dealers []models.Dealer
	if err := json.NewDecoder(rr.Body).Decode(&dealers); err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	if len(dealers) != 2 {
		t.Errorf("expected 2 dealers, got %v", len(dealers))
	}

	if dealers[0].Name != "Dealer Test A" {
		t.Errorf("expected dealer name 'Dealer Test A', got %v", dealers[0].Name)
	}

	// Verify all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
