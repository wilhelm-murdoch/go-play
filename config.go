package play

import "fmt"

var (
	ConfigDefault    = NewConfig("https://go.dev", "_/share", "play/p")
	ConfigDeprecated = NewConfig("http://play.golang.org", "share", "p")
)

// Config
type Config struct {
	urlBase  string
	urlPost  string
	urlShare string
}

// NewConfig
func NewConfig(base, post, share string) Config {
	return Config{
		urlBase:  base,
		urlPost:  post,
		urlShare: share,
	}
}

// GetPostUrl
func (c Config) GetPostUrl() string {
	return fmt.Sprintf("%s/%s", c.urlBase, c.urlPost)
}

// GetShareUrl
func (c Config) GetShareUrl() string {
	return fmt.Sprintf("%s/%s", c.urlBase, c.urlShare)
}
