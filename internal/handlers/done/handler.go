package done

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/GalichAnton/go_final_project/internal/services"
	"github.com/GalichAnton/go_final_project/internal/utils"
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
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.ErrorResponse(w, "Не указан идентификатор", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetById(id)
	if err != nil {
		utils.ErrorResponse(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	if task.Repeat == "" {
		err = h.service.Delete(id)
		if err != nil {
			utils.ErrorResponse(w, "Error deleting task: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		now := time.Now()
		task.Date, err = utils.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			utils.ErrorResponse(w, "правило повторения указано в неправильном формате", http.StatusBadRequest)
			return
		}
		err = h.service.Update(task)
		if err != nil {
			utils.ErrorResponse(w, "Error updating task: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{})
}
