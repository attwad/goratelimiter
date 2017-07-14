// Package limiter implements a simple blocking rate limiter.
package limiter

import "time"

// Rate limiter is a blocking rate limiter.
type RateLimiter interface {
	// Wait will block until the next request can be made.
	Wait()
}

type limiter struct {
	throttle <-chan time.Time
}

// NewRateLimiter creates a new rate limiter that limits to 1 query per the given duration.
// It will panic if duration <= 0
func NewRateLimiter(timeBetweenRequests time.Duration) RateLimiter {
	if timeBetweenRequests <= 0*time.Second {
		panic("time between requests need to be > 0")
	}
	return &limiter{
		throttle: time.Tick(timeBetweenRequests),
	}
}

// Wait will block until the next request can be made.
func (l *limiter) Wait() {
	<-l.throttle
}
