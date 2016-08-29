package scheduler

import (
	"time"
)

// TickerScheduler uses emits a tick every time.Duration as defined by the
// Duration field.
type TickerScheduler struct {
	Duration time.Duration
	ticker   <-chan time.Time
}

// Ticker emits a struct every time.Duration.
func (t *TickerScheduler) Ticker() <-chan struct{} {
	ticker := make(chan struct{}, 1)
	go func(output chan struct{}) {
		ticker := time.Tick(t.Duration)
		for {
			<-ticker
			output <- struct{}{}
		}
	}(ticker)

	return ticker
}

// Errors for a TickerScheduler returns a nil as there are no errors
// emitted by a time.Ticker.
func (t *TickerScheduler) Errors() <-chan error {
	return nil
}
