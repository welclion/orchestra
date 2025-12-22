// Пакет model содержит сущность Project.
package model

import "time"

// Project — ИТ-проект в системе ORCHESTRA.
type Project struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	CustomerID string    `json:"customer_id"`
	Status     string    `json:"status"` // draft | in_progress | completed
	CreatedAt  time.Time `json:"created_at"`
}
