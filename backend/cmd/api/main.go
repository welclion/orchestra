// Точка входа в приложение ORCHESTRA.
package main

import (
	"log"
	"net/http"

	"orchestra/backend/internal/config"
	"orchestra/backend/internal/identity/handler"
	"orchestra/backend/internal/identity/repository"
	"orchestra/backend/internal/identity/service"
	"orchestra/backend/pkg/db"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadDBConfig()

	// Подключаемся к БД
	database := db.Connect(cfg.DSN())
	defer database.Close()

	// Инициализируем слои Identity
	userRepo := repository.NewUserRepository(database)
	authService := service.NewAuthService(userRepo)

	// Регистрируем эндпоинты
	http.HandleFunc("/health", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok","service":"ORCHESTRA"}`))
	}))

	http.HandleFunc("/auth/register", enableCORS(handler.RegisterHandler(authService)))

	// Запускаем сервер
	log.Println("🚀 ORCHESTRA backend запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем только фронтенд
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

		// Обрабатываем preflight-запрос (OPTIONS)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next(w, r)
	}
}
