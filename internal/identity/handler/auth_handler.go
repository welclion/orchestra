// Пакет handler обрабатывает HTTP-запросы для аутентификации.
package handler

import (
	"encoding/json"
	"net/http"

	"orchestra-backend/internal/identity/service"
)

// RegisterRequest — структура входных данных для регистрации.
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterHandler возвращает HTTP-хендлер для регистрации.
func RegisterHandler(authService *service.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Читаем JSON из тела запроса
		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Неверный JSON", http.StatusBadRequest)
			return
		}

		// Вызываем сервис
		user, err := authService.Register(req.Email, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Возвращаем успешный ответ
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"id":         user.ID,
			"email":      user.Email,
			"role_type":  user.RoleType,
			"created_at": user.CreatedAt,
		})
	}
}
