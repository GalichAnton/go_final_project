package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func AtomicLevel(logLevel string) zap.AtomicLevel {
	var level zapcore.Level
	if err := level.Set(logLevel); err != nil {
		log.Fatalf("failed to set log level: %v", err)
	}

	return zap.NewAtomicLevelAt(level)
}
