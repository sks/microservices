package env

import (
	"os"
	"strings"
)

// Env denotes the environment in which the app is running
type Env int

const (
	// PRODUCTION ...
	PRODUCTION Env = iota
	// STAGING ...
	STAGING
	// TESTING ...
	TESTING
	// DEV ...
	DEV
)

// Environment read the environment
type Environment interface {
	GetEnv() Env
	GetAppName() string
	IsDebugEnabled() bool
	GetRequiredEnvVar(key string) (string, error)
	GetValOrDefault(key, defaultValue string) string
}

type env struct {
}

// NewEnvFx create a new env fx
func NewEnvFx() Environment {
	return &env{}
}

func (e *env) GetEnv() Env {
	env := os.Getenv("ENVIRONMENT")
	switch strings.ToLower(env) {
	case "prod":
		return PRODUCTION
	case "staging":
		return STAGING
	case "testing":
		return TESTING
	default:
		return DEV
	}
}

func (e *env) GetAppName() (app string) {
	app = os.Getenv("APP_NAME")
	if app == "" {
		return "unknown_app"
	}
	return app
}

func (e *env) IsDebugEnabled() bool {
	return os.Getenv("DEBUG") != ""
}

func (e *env) GetValOrDefault(key, defaultValue string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return defaultValue
}

func (e *env) GetRequiredEnvVar(key string) (string, error) {
	val := os.Getenv(key)
	if val != "" {
		return val, nil
	}
	return "", NewMissingRequiredEnv(key)
}
