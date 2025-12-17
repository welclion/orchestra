// Пакет repository отвечает за работу с базой данных для пользователей.
package repository

import (
	"database/sql"
	"orchestra-backend/internal/identity/model"
)

// UserRepository — репозиторий для работы с пользователями в PostgreSQL.
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository создаёт новый экземпляр UserRepository.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create сохраняет нового пользователя в БД и возвращает его данные.
func (r *UserRepository) Create(email, passwordHash string) (*model.User, error) {
	query := `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
		RETURNING id, email, role_type, created_at
	`

	user := &model.User{}
	err := r.db.QueryRow(query, email, passwordHash).Scan(
		&user.ID,
		&user.Email,
		&user.RoleType,
		&user.CreatedAt,
	)

	return user, err
}
