package service

import (
	"orchestra/backend/internal/project/model"
	"orchestra/backend/internal/project/repository"
)

type ProjectService struct {
	repo *repository.ProjectRepository
}

func NewProjectService(repo *repository.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) CreateProject(name, customerID string) (*model.Project, error) {
	if name == "" {
		return nil, nil // можно добавить ошибку позже
	}
	return s.repo.Create(name, customerID)
}

func (s *ProjectService) GetProjectsByCustomer(customerID string) ([]*model.Project, error) {
	return s.repo.FindByCustomerID(customerID)
}
