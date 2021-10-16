package model

type Activity struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	DueDate   string `json:"dueDate"`
	Completed bool   `json:"completed"`
}
