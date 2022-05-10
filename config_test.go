package play_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilhelm-murdoch/go-play"
)

func TestConfigDefaultPost(t *testing.T) {
	var (
		def      = play.ConfigDefault
		expected = "https://go.dev/_/share"
	)

	assert.Equal(t, def.GetPostUrl(), expected, "expected the returned url to be %s, but got %s instead", def.GetPostUrl(), expected)
}

func TestConfigDefaultShare(t *testing.T) {
	var (
		def      = play.ConfigDefault
		expected = "https://go.dev/play/p"
	)

	assert.Equal(t, def.GetShareUrl(), expected, "expected the returned url to be %s, but got %s instead", def.GetShareUrl(), expected)
}

func TestConfigDeprecatedPost(t *testing.T) {
	var (
		dep      = play.ConfigDeprecated
		expected = "http://play.golang.org/share"
	)

	assert.Equal(t, dep.GetPostUrl(), expected, "expected the returned url to be %s, but got %s instead", dep.GetPostUrl(), expected)
}

func TestConfigDeprecatedShare(t *testing.T) {
	var (
		dep      = play.ConfigDeprecated
		expected = "http://play.golang.org/p"
	)

	assert.Equal(t, dep.GetShareUrl(), expected, "expected the returned url to be %s, but got %s instead", dep.GetShareUrl(), expected)
}

func TestConfigCustomPost(t *testing.T) {
	var (
		dep      = play.NewConfig("https://does-not-work.io", "share", "play")
		expected = "https://does-not-work.io/share"
	)

	assert.Equal(t, dep.GetPostUrl(), expected, "expected the returned url to be %s, but got %s instead", dep.GetPostUrl(), expected)
}

func TestConfigCustomShare(t *testing.T) {
	var (
		dep      = play.NewConfig("https://does-not-work.io", "share", "play")
		expected = "https://does-not-work.io/play"
	)

	assert.Equal(t, dep.GetShareUrl(), expected, "expected the returned url to be %s, but got %s instead", dep.GetShareUrl(), expected)
}
