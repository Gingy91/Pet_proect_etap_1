package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"pet_project_etap_1/internal/database"
	"pet_project_etap_1/internal/handlers"
	"pet_project_etap_1/internal/taskService"
)

func main() {
	// Инициализация БД
	database.InitDB()
	err := database.DB.AutoMigrate(&taskService.Task{})
	if err != nil {
		return
	}

	// Создание слоёв
	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/update/{id}", handler.UpdateTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
