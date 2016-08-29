package birds

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

// Tweeter tweets on behalf of a particular set of access keys.
type Tweeter interface {
	// Post a tweet as this Tweeter
	PostTweet(status string) error
}

// Anaconda uses github.com/ChimeraCoder/anaconda to post tweets via
// the Twitter API.
type Anaconda struct {
	api *anaconda.TwitterApi
}

// NewAnaconda creates and initializes a new Anaconda API client.
func NewAnaconda(accessToken string, secretToken string) *Anaconda {
	return &Anaconda{
		api: anaconda.NewTwitterApi(accessToken, secretToken),
	}
}

// PostTweet - See Tweeter.PostTweet
func (a *Anaconda) PostTweet(status string) (err error) {
	_, err = a.api.PostTweet(status, url.Values{})
	return err
}
