package service

import (
	"database/sql"
	"errors"
	"strings"

	"orchestra/backend/internal/identity/model"
	"orchestra/backend/internal/identity/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(email, password string) (*model.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email и пароль обязательны")
	}
	if !strings.Contains(email, "@") {
		return nil, errors.New("некорректный email")
	}
	if len(password) < 6 {
		return nil, errors.New("пароль должен быть не короче 6 символов")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("ошибка хэширования")
	}

	return s.repo.Create(email, string(hash))
}

func (s *AuthService) Login(email, password string) (*model.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email и пароль обязательны")
	}

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("неверный email или пароль")
		}
		return nil, errors.New("ошибка сервера")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("неверный email или пароль")
	}

	return user, nil
}
