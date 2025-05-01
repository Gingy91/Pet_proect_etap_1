package main

import (
	"github.com/gorilla/mux"
	"pet_project_etap_1/internal/database"
	"pet_project_etap_1/internal/handlers"
	"pet_project_etap_1/internal/taskService"
)

func main() {
	// Инициализация БД
	database.InitDB()

	// Создание слоёв
	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/tasks/get", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/v1/tasks/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/v1/tasks/patch/{id}", handler.UpdateTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/v1/tasks/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")

}
