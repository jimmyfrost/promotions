package promoDBService

import (
	"database/sql"
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
		shorthand TEXT NOT NULL,
		name TEXT NOT NULL,
		title TEXT NOT NULL
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

	log.Println("Database initialized successfully!")
}

func CreateNewEmployee(shorthand, name, title string) (int64, error) {
	var id int64
	err := DB.QueryRow(
		"INSERT INTO employees (shorthand, name, title) VALUES ($1, $2, $3) RETURNING id",
		shorthand, name, title,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func DeleteEmployee(id int64) error {
	_, err := DB.Exec("DELETE FROM employees WHERE id = $1", id)
	return err
}

func ReadEmployeesList() []models.Employee {
	rows, err := DB.Query("SELECT id, shorthand, name, title FROM employees")
	if err != nil {
		log.Println("Error reading employees:", err)
		return nil
	}
	defer rows.Close()

	emps := make([]models.Employee, 0)
	for rows.Next() {
		var emp models.Employee
		err := rows.Scan(&emp.ID, &emp.Shorthand, &emp.Name, &emp.Title)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		emps = append(emps, emp)
	}

	return emps
}
