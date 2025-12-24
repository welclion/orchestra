package service

import (
	"orchestra/backend/internal/workflow/model"
	"orchestra/backend/internal/workflow/repository"
)

type WorkflowService struct {
	repo *repository.WorkflowRepository
}

func NewWorkflowService(repo *repository.WorkflowRepository) *WorkflowService {
	return &WorkflowService{repo: repo}
}

func (s *WorkflowService) CreateStage(name string, projectID string, order int) (*model.Stage, error) {
	if name == "" || projectID == "" {
		return nil, nil // можно добавить ошибку
	}
	return s.repo.CreateStage(name, projectID, order)
}

func (s *WorkflowService) CreateRole(name string, stageID string) (*model.Role, error) {
	if name == "" || stageID == "" {
		return nil, nil
	}
	return s.repo.CreateRole(name, stageID)
}
