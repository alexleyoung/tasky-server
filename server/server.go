package server

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	initEndpoints()

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initEndpoints() {
	http.HandleFunc("/tasks", taskHandlers)
}