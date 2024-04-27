package task

import "github.com/GalichAnton/go_final_project/internal/models/task"

func (s *service) Update(task *task.Task) error {
	err := s.taskRepository.Update(task)
	if err != nil {
		return err
	}

	return nil
}
