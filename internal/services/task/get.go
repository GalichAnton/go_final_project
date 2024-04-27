package task

import (
	"github.com/GalichAnton/go_final_project/internal/models/task"
)

func (s *service) GetById(id string) (*task.Task, error) {
	task, err := s.taskRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}
