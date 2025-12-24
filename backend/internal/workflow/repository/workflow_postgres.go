package repository

import (
	"database/sql"
	"orchestra/backend/internal/workflow/model"
)

type WorkflowRepository struct {
	db *sql.DB
}

func NewWorkflowRepository(db *sql.DB) *WorkflowRepository {
	return &WorkflowRepository{db: db}
}

// CreateStage создаёт новый этап
func (r *WorkflowRepository) CreateStage(name string, projectID string, order int) (*model.Stage, error) {
	query := `
		INSERT INTO stages (name, project_id, "order")
		VALUES ($1, $2, $3)
		RETURNING id, name, project_id, "order", status, created_at
	`
	stage := &model.Stage{}
	err := r.db.QueryRow(query, name, projectID, order).Scan(
		&stage.ID, &stage.Name, &stage.ProjectID, &stage.Order,
		&stage.Status, &stage.CreatedAt,
	)
	return stage, err
}

// CreateRole создаёт роль на этапе
func (r *WorkflowRepository) CreateRole(name string, stageID string) (*model.Role, error) {
	query := `
		INSERT INTO roles (name, stage_id)
		VALUES ($1, $2)
		RETURNING id, name, stage_id, user_id, status, created_at
	`
	role := &model.Role{}
	err := r.db.QueryRow(query, name, stageID).Scan(
		&role.ID, &role.Name, &role.StageID, &role.UserID,
		&role.Status, &role.CreatedAt,
	)
	return role, err
}
