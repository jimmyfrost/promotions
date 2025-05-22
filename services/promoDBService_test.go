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

func TestDeleteEmployee(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing sqlmock: %v", err)
	}
	defer mockDB.Close()
	DB = mockDB

	mock.ExpectExec("DELETE FROM employees").WithArgs(int64(1)).WillReturnResult(sqlmock.NewResult(0, 1))
	err = DeleteEmployee(1)
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestReadEmployeesList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing sqlmock: %v", err)
	}
	defer mockDB.Close()
	DB = mockDB

	rows := sqlmock.NewRows([]string{"id", "email", "name", "title", "track"}).
		AddRow(1, "test@example.com", "John Doe", "Engineer", "Engineering")
	mock.ExpectQuery("SELECT id, email, name, title, track FROM employees").WillReturnRows(rows)

	emps := ReadEmployeesList()
	assert.Len(t, emps, 1)
	assert.Equal(t, int64(1), emps[0].ID)
	assert.Equal(t, "test@example.com", emps[0].Email)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestGetAchievementsByEmployeeID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing sqlmock: %v", err)
	}
	defer mockDB.Close()
	DB = mockDB

	rows := sqlmock.NewRows([]string{"situation", "task", "action", "result"}).
		AddRow("Situation1", "Task1", "Action1", "Result1")
	mock.ExpectQuery("SELECT situation, task, action, result FROM achievements WHERE employee_id = ?").WithArgs(int64(1)).WillReturnRows(rows)

	achievements := GetAchievementsByEmployeeID(1)
	assert.Len(t, achievements, 1)
	assert.Equal(t, "Situation1", achievements[0].Situation)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestCreateGoal(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing sqlmock: %v", err)
	}
	defer mockDB.Close()
	DB = mockDB

	mock.ExpectExec("INSERT INTO goals").
		WithArgs("GoalTitle", "GoalDetails", 12, int64(1)).
		WillReturnResult(sqlmock.NewResult(2, 1))

	id, err := CreateGoal("GoalTitle", "GoalDetails", 12, 1)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), id)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestCreateSuggestion(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing sqlmock: %v", err)
	}
	defer mockDB.Close()
	DB = mockDB

	mock.ExpectExec("INSERT INTO goals").
		WithArgs("1", "SuggestionTitle", "SuggestionDetails", int64(1)).
		WillReturnResult(sqlmock.NewResult(3, 1))

	id, err := CreateSuggestion("1", "SuggestionTitle", "SuggestionDetails", 1)
	assert.NoError(t, err)
	assert.Equal(t, int64(3), id)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestCreateAchievement(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing sqlmock: %v", err)
	}
	defer mockDB.Close()
	DB = mockDB

	mock.ExpectExec("INSERT INTO goals").
		WithArgs("Situation", "Task", "Action", "Result", int64(1)).
		WillReturnResult(sqlmock.NewResult(4, 1))

	id, err := CreateAchievement("Situation", "Task", "Action", "Result", 1)
	assert.NoError(t, err)
	assert.Equal(t, int64(4), id)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestDeleteAchievement(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing sqlmock: %v", err)
	}
	defer mockDB.Close()
	DB = mockDB

	mock.ExpectExec("DELETE FROM achievements").WithArgs(int64(1)).WillReturnResult(sqlmock.NewResult(0, 1))

	err = DeleteAchievement(1)
	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
