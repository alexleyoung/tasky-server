package db

type Task struct {
	ID        int
	Name      string
	Description string
	DueDate   string
	Completed bool
}