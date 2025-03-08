package models

type Employee struct {
	ID        int64  `json:"empId"`
	Shorthand string `json:"shorthand"`
	Name      string `json:"name"`
	Title     string `json:"empTitle"`
	Goals     []Goal `json:"goals"`
}

type Goal struct {
	ID                  int64        `json:"goalId"`
	Title               string       `json:"goalTitle"`
	Details             string       `json:"goalDetails"`
	TimeHorizonInMonths int          `json:"timeHorizon"`
	EmployeeID          int64        `json:"fkEmpId"`
	Suggestions         []Suggestion `json:"suggestions"`
}

type Suggestion struct {
	ID      int64  `json:"suggId"`
	Type    int8   `json:"type"`
	Title   string `json:"suggTitle"`
	Details string `json:"suggDetails"`
	GoalID  int64  `json:"fkGoalId"`
}
