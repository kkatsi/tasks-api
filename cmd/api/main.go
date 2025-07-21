package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"rest-api/internal/handler"
	"rest-api/internal/routes"
	"rest-api/internal/service"
	"rest-api/internal/storage"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	store := storage.NewSQLiteStore(db)
	taskService := service.NewTaskService(store)
	handler := handler.NewTaskHandler(taskService)

	mux := http.NewServeMux()

	routes.SetupRoutes(mux, handler)

	// Start server
	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
