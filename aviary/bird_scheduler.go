package aviary

import (
	"github.com/grepory/birdin/birds"
	"github.com/grepory/birdin/scheduler"
)

// BirdScheduler is a combination of a Bird and a Scheduler for use in the
// Aviary.
type BirdScheduler struct {
	Bird      birds.Bird
	Scheduler scheduler.Scheduler
}

// Start a BirdScheduler. Errors from either Bird or Scheduler are emitted
// over the returned error channel, and tweets are emitted over the returned
// string channel.
func (bs *BirdScheduler) Start() (<-chan string, <-chan error) {
	errs := make(chan error, 1)
	tweets := make(chan string, 1)
	go func() {
		defer close(errs)
		defer close(tweets)
		for {
			select {
			case <-bs.Scheduler.Ticker():
				tweet, err := bs.Bird.Tweet()
				if err != nil {
					errs <- err
				}
				tweets <- tweet
			case err := <-bs.Scheduler.Errors():
				errs <- err
			}
		}
	}()
	return tweets, errs
}
