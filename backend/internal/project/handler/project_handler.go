package handler

import (
	"encoding/json"
	"net/http"
	"orchestra/backend/internal/project/service"
)

type CreateProjectRequest struct {
	Name string `json:"name"`
}

func CreateProjectHandler(projectService *service.ProjectService, customerID string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateProjectRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Неверный JSON", http.StatusBadRequest)
			return
		}

		project, err := projectService.CreateProject(req.Name, customerID)
		if err != nil {
			http.Error(w, "Ошибка создания проекта", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(project)
	}
}

func GetProjectsHandler(projectService *service.ProjectService, customerID string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projects, err := projectService.GetProjectsByCustomer(customerID)
		if err != nil {
			http.Error(w, "Ошибка загрузки проектов", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"projects": projects,
		})
	}
}
