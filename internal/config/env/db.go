package env

import (
	"os"

	"github.com/pkg/errors"
)

const (
	DBPathEnvName = "TODO_DBFILE"
)

type DBConfig interface {
	Path() string
}

type dbConfig struct {
	path string
}

func NewDBConfig() (DBConfig, error) {
	path := os.Getenv(DBPathEnvName)
	if len(path) == 0 {
		return nil, errors.New("db file not found")
	}

	return &dbConfig{
		path: path,
	}, nil
}

func (cfg *dbConfig) Path() string {
	return cfg.path
}
