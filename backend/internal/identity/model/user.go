// internal/identity/model/user.go
package model

import "time"

// User — модель пользователя в системе ORCHESTRA.
type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // скрыт в JSON
	RoleType     string    `json:"role_type"`
	CreatedAt    time.Time `json:"created_at"`
}
