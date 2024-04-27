package task

import (
	modelService "github.com/GalichAnton/go_final_project/internal/models/task"
)

func (s *service) Create(info *modelService.Info) (int64, error) {
	id, err := s.taskRepository.Create(info)
	if err != nil {
		return 0, err
	}

	return id, nil
}
