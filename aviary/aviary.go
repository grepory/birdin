// Package aviary provides a management layer for all of our tweet bots.
// Tweet bots (a BirdScheduler) are made up of a tweet generator (a Bird)
// and a scheduler (a Scheduler). An Aviary is simply a collection of these
// with a simple startup and logging mechanism.
package aviary

import (
	"sync"

	"github.com/Sirupsen/logrus"
)

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
