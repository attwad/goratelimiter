package limiter

import (
	"testing"
	"time"
)

func TestNegativeDuration(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Negative duration between requests did not panic")
		}
	}()
	_ = NewRateLimiter(-1 * time.Nanosecond)
}

func TestBlocks(t *testing.T) {
	l := NewRateLimiter(1 * time.Nanosecond)
	l.Wait()
}
