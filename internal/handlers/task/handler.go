package task

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/GalichAnton/go_final_project/internal/models/task"
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
	switch r.Method {
	case http.MethodGet:
		h.handleGet(w, r)
	case http.MethodPost:
		h.handlePost(w, r)
	case http.MethodPut:
		h.handlePut(w, r)
	case http.MethodDelete:
		h.handleDelete(w, r)
	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) handleGet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.ErrorResponse(w, "Не указан идентификатор", http.StatusBadRequest)
		return
	}

	dbTask, err := h.service.GetById(id)
	if err != nil {
		utils.ErrorResponse(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(dbTask)

}

func (h *Handler) handlePost(w http.ResponseWriter, r *http.Request) {
	var taskInfo task.Info

	err := json.NewDecoder(r.Body).Decode(&taskInfo)
	if err != nil {
		utils.ErrorResponse(w, "ошибка десериализации JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if taskInfo.Title == "" {
		utils.ErrorResponse(w, "Не указан заголовок задачи", http.StatusBadRequest)
		return
	}

	now := time.Now()
	startDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if taskInfo.Date == "" {
		taskInfo.Date = now.Format("20060102")
	}

	date, err := time.Parse("20060102", taskInfo.Date)
	if err != nil {
		utils.ErrorResponse(w, "дата представлена в формате, отличном от 20060102", http.StatusBadRequest)
		return
	}

	if date.Before(startDay) {
		if taskInfo.Repeat == "" {
			taskInfo.Date = now.Format("20060102")
		} else {
			taskInfo.Date, err = utils.NextDate(now, taskInfo.Date, taskInfo.Repeat)
			if err != nil {
				utils.ErrorResponse(w, "правило повторения указано в неправильном формате", http.StatusBadRequest)
				return
			}
		}
	}

	var response task.Response
	response.ID, err = h.service.Create(&taskInfo)
	if err != nil {
		utils.ErrorResponse(w, "Error creating task: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) handlePut(w http.ResponseWriter, r *http.Request) {
	var newTask task.Task

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		utils.ErrorResponse(w, "ошибка десериализации JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if newTask.ID == "" {
		utils.ErrorResponse(w, "Не указан ID задачи", http.StatusBadRequest)
		return
	}

	_, err = strconv.Atoi(newTask.ID)
	if err != nil {
		utils.ErrorResponse(w, "ID задачи не является числом", http.StatusBadRequest)
		return
	}

	if newTask.Title == "" {
		utils.ErrorResponse(w, "Не указан заголовок задачи", http.StatusBadRequest)
		return
	}

	now := time.Now()
	startDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if newTask.Date == "" {
		newTask.Date = now.Format("20060102")
	}

	date, err := time.Parse("20060102", newTask.Date)
	if err != nil {
		utils.ErrorResponse(w, "дата представлена в формате, отличном от 20060102", http.StatusBadRequest)
		return
	}

	if date.Before(startDay) {
		if newTask.Repeat == "" {
			newTask.Date = now.Format("20060102")
		} else {
			newTask.Date, err = utils.NextDate(now, newTask.Date, newTask.Repeat)
			if err != nil {
				utils.ErrorResponse(w, "правило повторения указано в неправильном формате", http.StatusBadRequest)
				return
			}
		}
	}

	err = h.service.Update(&newTask)
	if err != nil {
		utils.ErrorResponse(w, "Error updating newTask: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{})
}

func (h *Handler) handleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.ErrorResponse(w, "Не указан идентификатор", http.StatusBadRequest)
		return
	}

	err := h.service.Delete(id)
	if err != nil {
		utils.ErrorResponse(w, "Ошибка при удалении", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{})
}
