package play

import "fmt"

var (
	ConfigDefault    = NewConfig("https://go.dev", "_/share", "play/p")  // A configuration representing default values for the current Go Playground.
	ConfigDeprecated = NewConfig("http://play.golang.org", "share", "p") // A configuration representing the deprecated Go Playground.
)

// Config exposes URL patterns used to interact with various versions of the
// Go Playground.
type Config struct {
	urlBase  string // The base URL for the target playground.
	urlPost  string // Submission endpoint suffix for the base URL.
	urlShare string // Play endpoint for the base URL.
}

// NewConfig returns a new instance of struct Config using the specified
// portions of the target Go Playground endpoints.
func NewConfig(base, post, share string) Config {
	return Config{
		urlBase:  base,
		urlPost:  post,
		urlShare: share,
	}
}

// GetPostUrl returns a formatted string representing the full submission URL.
func (c Config) GetPostUrl() string {
	return fmt.Sprintf("%s/%s", c.urlBase, c.urlPost)
}

// GetShareUrl returns a formatted string representing the shareable play URL.
func (c Config) GetShareUrl() string {
	return fmt.Sprintf("%s/%s", c.urlBase, c.urlShare)
}
