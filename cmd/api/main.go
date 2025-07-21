package main

import (
	"fmt"
	"log"
	"net/http"
	"rest-api/internal/handler"
	"rest-api/internal/routes"
	"rest-api/internal/service"
	"rest-api/internal/storage"
)

func main() {

	store := storage.NewMemoryStore()
	taskService := service.NewTaskService(store)
	handler := handler.NewTaskHandler(taskService)

	mux := http.NewServeMux()

	routes.SetupRoutes(mux, handler)

	// Start server
	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
