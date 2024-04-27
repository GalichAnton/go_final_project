package task

import "github.com/GalichAnton/go_final_project/internal/models/task"

func (s *service) GetTasks(searchStr string) ([]task.Task, error) {
	tasks, err := s.taskRepository.GetTasks(searchStr)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
