package utils

import (
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWritePlaylistsItemsLinks(t *testing.T) {
	if runtime.GOOS == "windows" {
		// TODO: add windows tests

		return
	}

	WriteRetroarchPlaylist(RetroArchPlaylist{}, "./fake-exe")

	WritePlaylistsItemsLinks(
		[]PlaylistInfo{
			{
				Path: "/path/to/the/playlist",

				Content: RetroArchPlaylist{Items: []RetroArchPlaylistItem{
					{
						RomPath: "./rom-test",

						Label: "rom test",

						CorePath: "/path/to/core",
					},
				}},
			},
		},

		"./fake-exe",

		".",
	)

	s, _ := os.Stat("./rom test")

	assert.Greater(
		t,

		s.Size(),

		int64(0),

		"Should be able to write a playlist link",
	)

	os.Remove("./rom test")

	os.Remove("./fake-exe")
}
