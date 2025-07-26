package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"rest-api/internal/config"
	"rest-api/internal/handler"
	"rest-api/internal/routes"
	"rest-api/internal/service"
	"rest-api/internal/storage"
	"rest-api/internal/utils"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	config.InitConfig()

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	store := storage.NewSQLiteStore(db)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		utils.JsonResponse(w, http.StatusOK, map[string]string{"status": "OK"})
	})

	//tasks
	taskService := service.NewTaskService(store)
	taskHandler := handler.NewTaskHandler(taskService)
	routes.SetupTaskRoutes(mux, taskHandler)

	//users
	userService := service.NewUserService(store)
	userHandler := handler.NewUserHandler(userService)
	routes.SetupUserRoutes(mux, userHandler)

	//auth
	authService := service.NewAuthService(store, userService)
	authHandler := handler.NewAuthHandler(authService)
	routes.SetupAuthRoutes(mux, authHandler)

	// Start server
	port := config.Get().Port
	addr := ":" + strconv.Itoa(port)
	fmt.Printf("Server starting on port %d...\n", port)
	log.Fatal(http.ListenAndServe(addr, mux))
}
