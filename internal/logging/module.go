package logging

import (
	"os"
	"path/filepath"

	"github.com/sks/microservices/internal/env"
	"go.uber.org/zap"
)

// NewLogger create a new logging module
func NewLogger(envfx env.Environment) (logger *zap.Logger, err error) {
	logger, err = zap.NewProduction()
	if err != nil || envfx.GetEnv() == env.DEV || envfx.IsDebugEnabled() {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		return nil, err
	}
	return logger.With(
		zap.String("app_name", envfx.GetAppName()),
		zap.String("command", filepath.Base(os.Args[0])),
	), nil
}
