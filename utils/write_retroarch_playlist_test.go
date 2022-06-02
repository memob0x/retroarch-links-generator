package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteRetroarchPlaylist(t *testing.T) {
	f, _, _ := WriteRetroarchPlaylist(RetroArchPlaylist{}, "./playlist.lpl")

	assert.Equal(
		t,

		f.Name(),

		"./playlist.lpl",

		"Should be able to create a playlist file",
	)

	s, _ := os.Stat("./playlist.lpl")

	assert.Greater(
		t,

		s.Size(),

		int64(0),

		"Should be able to write a playlist file contents",
	)

	os.Remove("./playlist.lpl")
}
