package app

import (
	"log"

	"github.com/GalichAnton/go_final_project/internal/clients/db"
	"github.com/GalichAnton/go_final_project/internal/config"
	"github.com/GalichAnton/go_final_project/internal/config/env"
	"github.com/GalichAnton/go_final_project/internal/repository"
	"github.com/GalichAnton/go_final_project/internal/repository/task"
	"github.com/GalichAnton/go_final_project/internal/services"
	taskService "github.com/GalichAnton/go_final_project/internal/services/task"
)

type serviceProvider struct {
	dbConfig   config.DBConfig
	httpConfig config.HTTPConfig
	logConfig  config.LogConfig

	taskRepository repository.TaskRepository

	taskService services.TaskService
	pass        string

	dbClient db.Client
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) DBConfig() config.DBConfig {
	if s.dbConfig == nil {
		cfg, err := env.NewDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		s.dbConfig = cfg
	}

	return s.dbConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) LogConfig() config.LogConfig {
	if s.logConfig == nil {
		cfg, err := env.NewLogConfig()
		if err != nil {
			log.Fatalf("failed to get log config: %s", err.Error())
		}

		s.logConfig = cfg
	}

	return s.logConfig
}

func (s *serviceProvider) GetPassFromEnv() config.HTTPConfig {
	if s.pass == "" {
		cfg, err := env.NewPassConfig()
		if err != nil {
			log.Fatalf("failed to get password: %s", err.Error())
		}

		s.pass = cfg.GetPass()
	}

	return s.httpConfig
}

func (s *serviceProvider) DBClient() db.Client {
	if s.dbClient == nil {
		sqliteDb, err := db.New(s.DBConfig().Path())
		if err != nil {
			log.Fatalf("Failed to create db client: %v", err)
		}

		if err != nil {
			log.Fatalf("Ping error: %s", err.Error())
		}

		s.dbClient = sqliteDb
	}

	return s.dbClient
}

func (s *serviceProvider) TaskRepository() repository.TaskRepository {
	if s.taskRepository == nil {
		s.taskRepository = task.NewTaskRepository(s.DBClient())
	}

	return s.taskRepository
}

func (s *serviceProvider) TaskService() services.TaskService {
	if s.taskService == nil {
		s.taskService = taskService.NewService(
			s.TaskRepository(),
		)
	}

	return s.taskService
}
