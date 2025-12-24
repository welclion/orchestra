package model

import "time"

type Stage struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	ProjectID string    `json:"project_id"`
	Order     int       `json:"order"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
