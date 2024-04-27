package config

import (
	"github.com/joho/godotenv"
)

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

type HTTPConfig interface {
	Address() string
}

type DBConfig interface {
	Path() string
}

type Password interface {
	GetPass() string
}

type LogConfig interface {
	Level() string
}
