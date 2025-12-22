package repository

import (
	"database/sql"
	"orchestra/backend/internal/project/model"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Create(name, customerID string) (*model.Project, error) {
	query := `
		INSERT INTO projects (name, customer_id)
		VALUES ($1, $2)
		RETURNING id, name, customer_id, status, created_at
	`
	project := &model.Project{}
	err := r.db.QueryRow(query, name, customerID).Scan(
		&project.ID,
		&project.Name,
		&project.CustomerID,
		&project.Status,
		&project.CreatedAt,
	)
	return project, err
}

func (r *ProjectRepository) FindByCustomerID(customerID string) ([]*model.Project, error) {
	query := `SELECT id, name, customer_id, status, created_at FROM projects WHERE customer_id = $1`
	rows, err := r.db.Query(query, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*model.Project
	for rows.Next() {
		p := &model.Project{}
		err := rows.Scan(&p.ID, &p.Name, &p.CustomerID, &p.Status, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, nil
}
