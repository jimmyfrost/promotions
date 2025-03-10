package promoDBService

import (
	"database/sql"
	"fmt"
	"log"
	"promotions/internal/models"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDatabase() {
	var err error

	DB, err = sql.Open("sqlite", "file:promotions.db?mode=rwc")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS employees (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		name TEXT NOT NULL,
		title TEXT NOT NULL,
		track TEXT NOT NULL
	);`)
	if err != nil {
		log.Fatal("Error creating employees table:", err)
	}

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS goals (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		details TEXT,
		time_horizon_in_months INTEGER,
		employee_id INTEGER,
		FOREIGN KEY (employee_id) REFERENCES employees (id) ON DELETE CASCADE
	);`)
	if err != nil {
		log.Fatal("Error creating goals table:", err)
	}

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS suggestions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		type INTEGER NOT NULL,
		title TEXT NOT NULL,
		details TEXT,
		goal_id INTEGER,
		FOREIGN KEY (goal_id) REFERENCES goals (id) ON DELETE CASCADE
	);`)
	if err != nil {
		log.Fatal("Error creating suggestions table:", err)
	}

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS achievements (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		situation INTEGER NOT NULL,
		task TEXT NOT NULL,
		action TEXT NOT NULL,
		result TEXT NOT NULL,
		employee_id INTEGER,
		FOREIGN KEY (employee_id) REFERENCES employees (id) ON DELETE CASCADE
	);`)
	if err != nil {
		log.Fatal("Error creating suggestions table:", err)
	}

	log.Println("Database initialized successfully!")
}

func CreateNewEmployee(email, name, title, track string) (int64, error) {
	result, err := DB.Exec(
		"INSERT INTO employees (email, name, title, track) VALUES (?, ?, ?, ?)",
		email, name, title, track,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func DeleteEmployee(id int64) error {
	_, err := DB.Exec("DELETE FROM employees WHERE id = ?", id)
	fmt.Print("in delete...")
	return err
}

func ReadEmployeesList() []models.Employee {
	rows, err := DB.Query("SELECT id, email, name, title, track FROM employees")
	if err != nil {
		log.Println("Error reading employees:", err)
		return nil
	}
	defer rows.Close()

	emps := make([]models.Employee, 0)
	for rows.Next() {
		var emp models.Employee
		err := rows.Scan(&emp.ID, &emp.Email, &emp.Name, &emp.Title, &emp.Track)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		emps = append(emps, emp)
	}

	return emps
}

func CreateGoal(title string, details string, time int, employeeId int64) (int64, error) {
	result, err := DB.Exec(
		"INSERT INTO goals (title, details, time_horizon_in_months, employee_id) VALUES (?, ?, ?, ?)",
		title, details, time, employeeId,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil

}

func CreateSuggestion(suggestionType, title, details string, goalId int64) (int64, error) {
	result, err := DB.Exec(
		"INSERT INTO goals (type, title, details, goal_id) VALUES (?, ?, ?, ?)",
		suggestionType, title, details, goalId,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func CreateAchievement(situation, task, action, actionResult string, employeeId int64) (int64, error) {
	result, err := DB.Exec(
		"INSERT INTO goals (situation, title, details, result, employee_id) VALUES (?, ?, ?, ?, ?)",
		situation, task, action, actionResult, employeeId,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func DeleteAchievement(id int64) error {
	_, err := DB.Exec("DELETE FROM achievements WHERE id = ?", id)
	return err
}
