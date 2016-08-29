package birds

// A Bird can tweet.
type Bird interface {
	Name() string
	Tweet() (status string, err error)
}
