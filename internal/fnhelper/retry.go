package fnhelper

import (
	"fmt"
	"time"
)

// Retry retry the given function for maximum
func Retry(fn func() error, maxAttempts int, sleepDuration time.Duration) error {
	var err error
	for i := 0; i < maxAttempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		time.Sleep(sleepDuration)
	}
	return fmt.Errorf("Could not complete the call even after %d attempts, last error :%w", maxAttempts, err)
}
