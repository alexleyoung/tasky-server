package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexleyoung/taksy-server/db"
	_ "github.com/mattn/go-sqlite3"
)

func taskHandlers(w http.ResponseWriter, r *http.Request) {
	db := db.Connect()
	if r.Method == "GET" {
		getTasks(w, r, db)
	} else if r.Method == "POST" {
		createTask(w, r, db)
	} else if r.Method == "PUT" {
		updateTask(w, r, db)
	} else if r.Method == "DELETE" {
		deleteTask(w, r, db)
	}
}

func getTasks(w http.ResponseWriter, r *http.Request, c *sql.DB) {
	rows, err := c.Query("SELECT * FROM tasks")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	tasks := []db.Task{}

	for rows.Next() {
		var task db.Task
		err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.DueDate, &task.Completed)
		if err != nil {
			panic(err)
		}
		tasks = append(tasks, task)
	}

	fmt.Println(tasks)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request, c *sql.DB) {
	// Get json from body
	var task db.TaskPost
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate task
	if task.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}


	_, err := c.Exec("INSERT INTO tasks (name, description, due_date, completed) VALUES (?, ?, ?, ?)", task.Name, task.Description, task.DueDate, task.Completed)
	if err != nil {
		panic(err)
	}
}

func updateTask(w http.ResponseWriter, r *http.Request, c *sql.DB) {
	w.Write([]byte("Hello, world!"))
}

func deleteTask(w http.ResponseWriter, r *http.Request, c *sql.DB) {
	w.Write([]byte("Hello, world!"))
}
