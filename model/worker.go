package model

type Worker struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	DueDate   string `json:"dueDate"`
	Completed bool   `json:"completed"`
	Type      string `json:"type"`
}
