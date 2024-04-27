package task

func (s *service) Delete(id string) error {
	err := s.taskRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
