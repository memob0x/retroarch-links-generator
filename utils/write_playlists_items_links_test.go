package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testWritePlaylistItemsLinksWindows(_ *testing.T) {
	if !IS_OS_WINDOWS {
		return
	}

	// TODO: add windows tests
}

func testWritePlaylistItemsLinksNonWindows(t *testing.T) {
	if IS_OS_WINDOWS {
		return
	}

	WriteRetroarchPlaylist(RetroArchPlaylist{}, "./fake-exe")

	WritePlaylistsItemsLinks(
		[]PlaylistInfo{
			{
				Path: "/path/to/the/playlist",

				Content: RetroArchPlaylist{
					Items: []RetroArchPlaylistItem{
						{
							RomPath: "./rom-test",

							Label: "rom test",

							CorePath: "/path/to/core",
						},
					},
				},
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

func TestWritePlaylistsItemsLinks(t *testing.T) {
	testWritePlaylistItemsLinksWindows(t)

	testWritePlaylistItemsLinksNonWindows(t)
}
