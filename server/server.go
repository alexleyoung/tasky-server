package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexleyoung/taksy-server/db"
)

func Start() {
	db.InitDB()
	initEndpoints()

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initEndpoints() {
	http.HandleFunc("/tasks", taskHandlers)
}