package scheduler

// Scheduler emits messages at a configured interval over the Ticker()
// channel and emits any scheduler errors over the Errors() channel.
type Scheduler interface {
	Ticker() <-chan struct{}
	Errors() <-chan error
}
