package db

type Task struct {
	ID          int
	Name        string
	Description string
	DueDate     string
	Completed   bool
}

type TaskPost struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Completed   bool   `json:"completed"`
}