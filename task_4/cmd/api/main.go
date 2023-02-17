package main

import (
	"log"
	"net/http"

	"example.com/machine_test/task_4/pkg/db"
	"example.com/machine_test/task_4/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env : ", err)
	// }

	err := db.ConnectMongo()
	if err != nil {
		log.Fatal(err)
	}

	mux := mux.NewRouter()

	routes.InitializeRoutes(mux)

	log.Println("Starting server on port :3000")
	err = http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal("SERVER ERROR : ", err)
	}
}
