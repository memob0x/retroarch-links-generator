package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadShortcutsVdfFile(t *testing.T) {
	writtenShortcuts, _ := WritePlaylistsItemsVdf(
		[]PlaylistInfo{
			{
				Path: "/fake/path",

				Content: RetroArchPlaylist{
					Items: []RetroArchPlaylistItem{
						{
							RomPath: "a",

							Label: "b",

							CorePath: "c",
						},
					},
				},
			},
		},

		"../data/executable",

		"./read-playlists.vdf",
	)

	readShortcuts, _ := ReadShortcutsVdfFile("./read-playlists.vdf")

	assert.Equal(
		t,

		readShortcuts,

		writtenShortcuts,

		"Should be able to read Steam shortcuts from vdf file",
	)

	os.Remove("./read-playlists.vdf")
}
