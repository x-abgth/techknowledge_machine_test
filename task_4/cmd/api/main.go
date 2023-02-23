package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/machine_test/task_4/pkg/db"
	"example.com/machine_test/task_4/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var port string

func init() {
	// Loading the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env : ", err)
	}

	// Connecting to Mongo DB
	err = db.ConnectMongo()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Creating a new mux router
	mux := mux.NewRouter()

	routes.InitializeRoutes(mux)
	port = os.Getenv("SERVER_PORT")

	log.Printf("Starting server on port :%s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		log.Fatal("SERVER ERROR : ", err)
	}
}
