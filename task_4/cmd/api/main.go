package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/machine_test/task_4/pkg/routes"
	"github.com/gorilla/mux"
)

const (
	port = "3000"
)

func main() {
	// Creating a new mux router
	mux := mux.NewRouter()

	routes.InitializeRoutes(mux)

	log.Printf("Starting server on port :%s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		log.Fatal("SERVER ERROR : ", err)
	}
}
