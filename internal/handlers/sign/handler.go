package sign

import (
	"encoding/json"
	"net/http"
	"os"

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

type SignRequest struct {
	Password string `json:"password"`
}

type SignResponse struct {
	Token string `json:"token,omitempty"`
	Error string `json:"error,omitempty"`
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	var req SignRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.ErrorResponse(w, "Ошибка чтения пароля", http.StatusBadRequest)
		return
	}

	password := os.Getenv("TODO_PASSWORD")

	if req.Password != password {
		utils.ErrorResponse(w, "Неверный пароль", http.StatusBadRequest)
		return
	}

	tokenString := utils.CreateHash(req.Password)
	resp := SignResponse{Token: tokenString}
	json.NewEncoder(w).Encode(resp)
}
