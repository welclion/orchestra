package model

import "time"

type User struct {
	ID string 'json:"id"'
	Email string 'json:"email"'
	PasswordHash string 'json:"-"'
	RoleType string 'json:"role_type"'
	CreatedAt time.Time 'json:"created_at"'
}