package promoDBService

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewEmployee(t *testing.T) {
	// Create a new mock database
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing sqlmock: %v", err)
	}
	defer mockDB.Close()

	// Replace the global DB variable with our mock
	DB = mockDB

	// Set up expected SQL execution
	mock.ExpectExec("INSERT INTO employees").
		WithArgs("test@example.com", "John Doe", "Software Engineer", "Engineering").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the function being tested
	id, err := CreateNewEmployee("test@example.com", "John Doe", "Software Engineer", "Engineering")

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
