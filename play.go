package play

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

// Client represents an HTTP and Config instance combo used to make calls to the
// Go Playground.
type Client struct {
	http   http.Client // An HTTP client used to dispatch requests to the Go Playground
	config Config      // A configuration object used to compose Go Playground URLs
}

// NewClient returns a new instance of struct Client using the specified Config
// instance.
func NewClient(config Config) Client {
	client := http.Client{
		Timeout: 60 * time.Second,
	}

	return Client{
		http:   client,
		config: config,
	}
}

// FetchGroup takes a slice of byte slices representing multiple snippets of
// Go source code to process. This method makes use of the `errgroup` package
// which utilises Goroutines to process multiple snippets concurrently. This
// method stops on the first non-nil error response. The order of resulting
// slice of strings is not guaranteed.
func (c Client) FetchGroup(sources [][]byte) (result []string, err error) {
	errors := new(errgroup.Group)

	for _, source := range sources {
		source := source // shadow here as we cannot pass as parameter to error group func
		errors.Go(func() error {
			url, err := c.Fetch(source)
			if err != nil {
				return err
			}

			result = append(result, url)

			return nil
		})
	}

	if err := errors.Wait(); err != nil {
		return result, err
	}

	return result, err
}

// Fetch takes a single slice of bytes representing a snippet of Go source code
// to process in the Go Playground. This method returns either an error or a
// shareable Go Playground URL.
func (c Client) Fetch(source []byte) (result string, err error) {
	response, err := c.http.Post(c.config.GetPostUrl(), "raw", bytes.NewBuffer(source))
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	code, _ := ioutil.ReadAll(response.Body)
	return fmt.Sprintf("%s/%s", c.config.GetShareUrl(), string(code)), err
}
