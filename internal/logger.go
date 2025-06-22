package internal

import (
	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	return logger
}
