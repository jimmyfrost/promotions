package models

type Employee struct {
	ID          int64         `json:"empId"`
	Email       string        `json:"shorthand"`
	Name        string        `json:"name"`
	Title       string        `json:"empTitle"`
	Track       string        `json:"track"`
	Goals       []Goal        `json:"goals"`
	Achievement []Achievement `json:"achievement"`
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

type Achievement struct {
	ID         int64  `json:"achID"`
	Situation  string `json:"situation"`
	Task       string `json:"task"`
	Action     string `json:"action"`
	Result     string `json:"result"`
	ResumeLine string `json:"resumeLine"`
	EmployeeID int64  `json:"fkEmployeeId"`
}
