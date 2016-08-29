package aviary

import (
	"sync"

	"github.com/Sirupsen/logrus"
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

// An Aviary is a scheduler for a collection of Birds [birds.Bird].
type Aviary struct {
	birds []*BirdScheduler
}

// New creates a new Aviary.
func New(birds ...*BirdScheduler) *Aviary {
	return &Aviary{
		birds: birds,
	}
}

// Tweet starts all of the BirdSchedulers and logs messages from them all.
func (a *Aviary) Tweet() {
	wg := sync.WaitGroup{}
	for _, b := range a.birds {
		wg.Add(2)
		go func(bs *BirdScheduler) {
			logger := logrus.WithFields(logrus.Fields{
				"bird": bs.Bird.Name(),
			})
			tweetChan, errChan := bs.Start()
			logger.Info("BirdScheduler starting...")
			go func(tweets <-chan string) {
				for t := range tweets {
					logger.Info("Tweeting: ", t)
				}
				logger.Info("returning from tweet watcher")
				wg.Done()
			}(tweetChan)
			go func(errs <-chan error) {
				for e := range errs {
					logger.Error("Error tweeting: ", e)
				}
				logger.Info("returning from err watcher")
				wg.Done()
			}(errChan)
		}(b)
	}
	wg.Wait()
}
