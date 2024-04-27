package task

import (
	"github.com/GalichAnton/go_final_project/internal/repository"
	"github.com/GalichAnton/go_final_project/internal/services"
)

var _ services.TaskService = (*service)(nil)

type service struct {
	taskRepository repository.TaskRepository
}

func NewService(
	taskRepository repository.TaskRepository,
) *service {
	return &service{
		taskRepository: taskRepository,
	}
}
