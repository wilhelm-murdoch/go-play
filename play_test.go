package play_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilhelm-murdoch/go-play"
)

const playUrlPattern = `^https?:\/\/(www\.)?go\.dev\/play\/p\/\b[a-zA-Z0-9-_]{11,}$`

var (
	goCodes = [][]byte{
		[]byte(`package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}`),
		[]byte(`package main

import "fmt"

func main() {
	fmt.Println("Hello from another world!")
}`),
		[]byte(`abandon all hope ye who enter here`),
		[]byte(""),
	}
)

func TestFetch(t *testing.T) {
	p := play.NewClient(play.ConfigDefault)

	share, err := p.Fetch(goCodes[0])
	assert.Nil(t, err, "expected to return with no error, but got %s instead", err)
	assert.NotEmpty(t, share, "expected to return a valid url, but got nothing instead")
	assert.Regexp(t, playUrlPattern, share, "expected to return a valid url, but got %s instead", share)
}

func TestFetchGroup(t *testing.T) {
	p := play.NewClient(play.ConfigDefault)

	shares, err := p.FetchGroup(goCodes)
	assert.Equal(t, len(goCodes), len(shares), "expected to return %d valid urls, but got %d instead", len(goCodes), len(shares))
	assert.Nil(t, err, "expected to return with no error, but got %s instead", err)
	assert.NotEmpty(t, shares, "expected to return a valid url, but got nothing instead")
	for _, share := range shares {
		assert.Regexp(t, playUrlPattern, share, "expected to return a valid url, but got %s instead", share)
	}
}

func TestFetchWithCustomConfig(t *testing.T) {
	client := play.NewClient(play.NewConfig("https://does-not-work.io", "share", "play"))

	share, err := client.Fetch(goCodes[0])
	assert.NotNil(t, err, "expected to return with no error, but got %s instead", err)
	assert.Empty(t, share, "expected to return an empty url, but got %s instead", share)
}

func TestFetchGroupWithCustomConfig(t *testing.T) {
	client := play.NewClient(play.NewConfig("https://does-not-work.io", "share", "play"))

	shares, err := client.FetchGroup(goCodes)
	assert.NotEqual(t, len(goCodes), len(shares), "expected to return %d valid urls, but got %d instead", len(goCodes), len(shares))
	assert.NotNil(t, err, "expected to return with no error, but got %s instead", err)
	for _, share := range shares {
		assert.Empty(t, share, "expected to return an empty url, but got %s instead", share)
		assert.Regexp(t, playUrlPattern, share, "expected to return a valid url, but got %s instead", share)
	}
}
