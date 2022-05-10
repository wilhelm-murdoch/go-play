package play

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

// Client
type Client struct {
	http   http.Client
	config Config
}

// NewClient
func NewClient(config Config) Client {
	client := http.Client{
		Timeout: 60 * time.Second,
	}

	return Client{
		http:   client,
		config: config,
	}
}

// FetchGroup
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

// Fetch
func (c Client) Fetch(source []byte) (result string, err error) {
	client := http.Client{
		Timeout: 60 * time.Second,
	}

	response, err := client.Post(c.config.GetPostUrl(), "raw", bytes.NewBuffer(source))
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	code, _ := ioutil.ReadAll(response.Body)
	return fmt.Sprintf("%s/%s", c.config.GetShareUrl(), string(code)), err
}
