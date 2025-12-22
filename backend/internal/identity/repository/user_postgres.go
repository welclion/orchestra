package repository

import (
	"database/sql"

	"orchestra/backend/internal/identity/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(email, passwordHash string) (*model.User, error) {
	query := `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
		RETURNING id, email, role_type, created_at
	`
	user := &model.User{}
	err := r.db.QueryRow(query, email, passwordHash).Scan(
		&user.ID, &user.Email, &user.RoleType, &user.CreatedAt,
	)
	return user, err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	query := `SELECT id, email, password_hash, role_type, created_at FROM users WHERE email = $1`
	user := &model.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.RoleType, &user.CreatedAt,
	)
	return user, err
}
