package routes

import (
	"example.com/machine_test/task_4/internal/handlers"
	"github.com/gorilla/mux"
)

func InitializeRoutes(mux *mux.Router) {
	var app handlers.JsonResponse

	handlers.NewValidator()

	mux.HandleFunc("/items", app.GetAllItems).Methods("GET")
	mux.HandleFunc("/items/{id}", app.GetOneItem).Methods("GET")
	mux.HandleFunc("/items/{id}", app.CreateOneItem).Methods("POST")
	mux.HandleFunc("/items/{id}", app.UpdateOneItem).Methods("PUT")
	mux.HandleFunc("/items/{id}", app.DeleteOneItem).Methods("DELETE")
}
