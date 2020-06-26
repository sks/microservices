package env

import "fmt"

// ErrMissingRequiredEnv error thrown in case any required env variable is missing
type ErrMissingRequiredEnv struct {
	key string
}

func (e ErrMissingRequiredEnv) Error() string {
	return fmt.Sprintf("Missing required environment variable: %s", e.key)
}

// NewMissingRequiredEnv creates a new ErrMissingRequiredEnv
func NewMissingRequiredEnv(key string) ErrMissingRequiredEnv {
	return ErrMissingRequiredEnv{key}
}
