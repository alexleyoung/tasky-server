package server

import "net/http"

func taskHandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getTasks(w, r)
	} else if r.Method == "POST" {
		createTask(w, r)
	} else if r.Method == "PUT" {
		updateTask(w, r)
	} else if r.Method == "DELETE" {
		deleteTask(w, r)
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func createTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}