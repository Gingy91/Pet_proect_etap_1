package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"pet_project_etap_1/internal/database"
	"pet_project_etap_1/internal/handlers"
	"pet_project_etap_1/internal/taskService"
	"pet_project_etap_1/middleware"
)

func main() {
	// Инициализация БД
	database.InitDB()

	// Создание слоёв
	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	//Login нужен чтобы авторизовавынание работало
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	api := router.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.AuthMiddleware)
	api.HandleFunc("/tasks/get", handler.GetTasksHandler).Methods("GET")
	api.HandleFunc("/tasks/post", handler.PostTaskHandler).Methods("POST")
	api.HandleFunc("/tasks/patch/{id}", handler.UpdateTaskHandler).Methods("PATCH")
	api.HandleFunc("/tasks/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")

	fmt.Println("Сервер запускается ...")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Сервер запустился с ошибкой %v", err)
	}
}

//КОРОЧЕ ВСЕ СЛОМАЛ К ХУЯМ
