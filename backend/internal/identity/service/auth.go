// Пакет service содержит бизнес-логику для аутентификации.
package service

import (
	"errors"
	"strings"

	"orchestra/backend/internal/identity/model"
	"orchestra/backend/internal/identity/repository"

	"golang.org/x/crypto/bcrypt"
)

// AuthService — сервис для регистрации и входа.
type AuthService struct {
	repo *repository.UserRepository
}

// NewAuthService создаёт новый AuthService.
func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

// Register регистрирует нового пользователя.
func (s *AuthService) Register(email, password string) (*model.User, error) {
	// Простая валидация
	if email == "" || password == "" {
		return nil, errors.New("email и пароль обязательны")
	}
	if !strings.Contains(email, "@") {
		return nil, errors.New("некорректный email")
	}
	if len(password) < 6 {
		return nil, errors.New("пароль должен быть не короче 6 символов")
	}

	// Хэшируем пароль
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("ошибка хэширования пароля")
	}

	// Сохраняем в БД
	return s.repo.Create(email, string(passwordHash))
}
