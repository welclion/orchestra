// Точка входа в приложение ORCHESTRA.
package main

import (
	"log"
	"net/http"

	"orchestra-backend/internal/config"
	"orchestra-backend/internal/identity/handler"
	"orchestra-backend/internal/identity/repository"
	"orchestra-backend/internal/identity/service"
	"orchestra-backend/pkg/db"
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
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok","service":"ORCHESTRA"}`))
	})

	http.HandleFunc("/auth/register", handler.RegisterHandler(authService))

	// Запускаем сервер
	log.Println("🚀 ORCHESTRA backend запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
