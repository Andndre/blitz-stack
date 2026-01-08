package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"blitz-stack/backend/models"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetItems(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Mock data
	rows := sqlmock.NewRows([]string{"id", "title", "description"}).
		AddRow(1, "Item A", "Desc A").
		AddRow(2, "Item B", "Desc B")

	// Expect the query to be executed
	mock.ExpectQuery("SELECT id, title, description FROM items").WillReturnRows(rows)

	// Create a request
	req, err := http.NewRequest("GET", "/api/items", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := GetItems(db)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Assertions
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var items []models.Item
	if err := json.NewDecoder(rr.Body).Decode(&items); err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	if len(items) != 2 {
		t.Errorf("expected 2 items, got %v", len(items))
	}

	if items[0].Title != "Item A" {
		t.Errorf("expected item title 'Item A', got %v", items[0].Title)
	}

	// Verify all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}