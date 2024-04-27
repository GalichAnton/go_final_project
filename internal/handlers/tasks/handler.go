package tasks

import (
	"encoding/json"
	"net/http"

	"github.com/GalichAnton/go_final_project/internal/models/task"
	"github.com/GalichAnton/go_final_project/internal/services"
)

type Handler struct {
	service services.TaskService
}

func New(service services.TaskService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	searchStr := r.URL.Query().Get("search")
	tasks, err := h.service.GetTasks(searchStr)
	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	response := struct {
		Tasks []task.Task `json:"tasks"`
	}{Tasks: tasks}

	res, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
