package model

import "time"

type Role struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	StageID   string    `json:"stage_id"`
	UserID    *string   `json:"user_id,omitempty"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
