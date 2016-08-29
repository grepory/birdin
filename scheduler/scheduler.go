// Package scheduler provides schedulers for use with tweet bots (BirdScheduler).
// Currently there is only one scheduler implemented, the TickerScheduler which
// uses a time.Ticker to schedule tweets.
package scheduler

// Scheduler emits messages at a configured interval over the Ticker()
// channel and emits any scheduler errors over the Errors() channel.
type Scheduler interface {
	Ticker() <-chan struct{}
	Errors() <-chan error
}
