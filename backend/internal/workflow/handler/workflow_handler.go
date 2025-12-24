package handler

import (
	"encoding/json"
	"net/http"
	"orchestra/backend/internal/workflow/service"
)

type CreateStageRequest struct {
	Name      string `json:"name"`
	ProjectID string `json:"project_id"` // ← добавлено
	Order     int    `json:"order"`
}

func CreateStageHandler(workflowService *service.WorkflowService, userID string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateStageRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Неверный JSON", http.StatusBadRequest)
			return
		}

		// TODO: проверить, что проект принадлежит userID
		stage, err := workflowService.CreateStage(req.Name, req.ProjectID, req.Order)
		if err != nil {
			http.Error(w, "Ошибка создания этапа", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(stage)
	}
}
