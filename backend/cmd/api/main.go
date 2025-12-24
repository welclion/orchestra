package main

import (
	"log"
	"net/http"

	"orchestra/backend/internal/config"
	"orchestra/backend/internal/identity/handler"
	"orchestra/backend/internal/identity/repository"
	"orchestra/backend/internal/identity/service"
	projecthandler "orchestra/backend/internal/project/handler"
	projectrepo "orchestra/backend/internal/project/repository"
	projectservice "orchestra/backend/internal/project/service"
	workflowhandler "orchestra/backend/internal/workflow/handler"
	workflowrepo "orchestra/backend/internal/workflow/repository"
	workflowservice "orchestra/backend/internal/workflow/service"
	"orchestra/backend/pkg/db"
)

// Глобальный CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	cfg := config.LoadDBConfig()
	database := db.Connect(cfg.DSN())
	defer database.Close()

	// Identity
	userRepo := repository.NewUserRepository(database)
	authService := service.NewAuthService(userRepo)
	jwtService := service.NewJWTService("orchestra-secret-key-2025")

	// Project
	projectRepo := projectrepo.NewProjectRepository(database)
	projectService := projectservice.NewProjectService(projectRepo)

	// Workflow
	workflowRepo := workflowrepo.NewWorkflowRepository(database)
	workflowService := workflowservice.NewWorkflowService(workflowRepo)

	authMiddleware := handler.AuthMiddleware(jwtService)

	// Создаём маршрутизатор
	mux := http.NewServeMux()

	// Публичные маршруты
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})
	mux.HandleFunc("/auth/register", handler.RegisterHandler(authService))
	mux.HandleFunc("/auth/login", handler.LoginHandler(authService, jwtService))

	// Защищённый маршрут: GET/POST /projects
	mux.Handle("/projects", authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, ok := handler.GetUserFromContext(r.Context())
		if !ok {
			http.Error(w, "Пользователь не авторизован", http.StatusUnauthorized)
			return
		}
		switch r.Method {
		case http.MethodGet:
			projecthandler.GetProjectsHandler(projectService, userID).ServeHTTP(w, r)
		case http.MethodPost:
			projecthandler.CreateProjectHandler(projectService, userID).ServeHTTP(w, r)
		default:
			http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		}
	})))

	// Защищённый маршрут: POST /stages (временное решение)
	mux.Handle("/stages", authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Только POST", http.StatusMethodNotAllowed)
			return
		}
		userID, ok := handler.GetUserFromContext(r.Context())
		if !ok {
			http.Error(w, "Пользователь не авторизован", http.StatusUnauthorized)
			return
		}
		// Передаём userID и workflowService в хендлер
		workflowhandler.CreateStageHandler(workflowService, userID).ServeHTTP(w, r)
	})))

	// Оборачиваем всё в CORS
	finalHandler := corsMiddleware(mux)

	log.Println("🚀 ORCHESTRA backend запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", finalHandler))
}
