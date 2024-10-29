package server

import (
	"database/sql"
	"encoding/json"
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

func getTasks(w http.ResponseWriter, _ *http.Request, c *sql.DB) {
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// Expects a json body with the following fields:
// name, description, due_date, completed
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

// Expects a json body with the following fields:
// id, name, description, due_date, completed
func updateTask(w http.ResponseWriter, r *http.Request, c *sql.DB) {
	var task db.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate task
	if task.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	_, err := c.Exec("UPDATE tasks SET name = ?, description = ?, due_date = ?, completed = ? WHERE id = ?", task.Name, task.Description, task.DueDate, task.Completed, task.ID)
	if err != nil {
		panic(err)
	}
}

// Expects a query param with the id of the task to delete
func deleteTask(w http.ResponseWriter, r *http.Request, c *sql.DB) {
	// get id from request
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	res, err := c.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		panic(err)
	}

	if num, _ := res.RowsAffected(); num == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
}
