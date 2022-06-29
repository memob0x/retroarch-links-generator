package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRetroarchPlaylistsFiles(t *testing.T) {
	os.Mkdir("./playlists", 0755)

	WriteRetroarchPlaylist(RetroArchPlaylist{
		Items: []RetroArchPlaylistItem{
			{
				RomPath: "/path/to/rom",

				Label: "playlist 1 entry 1 label",

				CorePath: "/path/to/core",
			},
			{
				RomPath: "/path/to/rom",

				Label: "playlist 1 entry 2 label",

				CorePath: "/path/to/core",
			},
		},
	}, "./playlists/playlist-1.lpl")

	WriteRetroarchPlaylist(RetroArchPlaylist{
		Items: []RetroArchPlaylistItem{
			{
				RomPath: "/path/to/rom",

				Label: "playlist 2 entry 1 label",

				CorePath: "/path/to/core",
			},
			{
				RomPath: "/path/to/rom",

				Label: "playlist 2 entry 2 label",

				CorePath: "/path/to/core",
			},
		},
	}, "./playlists/playlist-2.lpl")

	WriteRetroarchPlaylist(RetroArchPlaylist{
		Items: []RetroArchPlaylistItem{
			{
				RomPath: "/path/to/rom",

				Label: "playlist 3 entry 1 label",

				CorePath: "/path/to/core",
			},
			{
				RomPath: "/path/to/rom",

				Label: "playlist 3 entry 2 label",

				CorePath: "/path/to/core",
			},
		},
	}, "./playlists/playlist-3.lpl")

	aS, _ := os.Stat("./playlists/playlist-1.lpl")

	bS, _ := os.Stat("./playlists/playlist-2.lpl")

	cS, _ := os.Stat("./playlists/playlist-3.lpl")

	parsedPlaylists, _ := ParseRetroarchPlaylistsFiles([]FileInfo{
		{
			Path:   "./playlists/playlist-1.lpl",
			FsFile: aS,
		},

		{
			Path:   "./playlists/playlist-2.lpl",
			FsFile: bS,
		},

		{
			Path:   "./playlists/playlist-3.lpl",
			FsFile: cS,
		},
	})

	assert.Equal(
		t,

		"playlist 1 entry 2 label",

		parsedPlaylists[0].Content.Items[1].Label,

		"should be able to parse playlist entry label",
	)

	assert.Equal(
		t,

		"playlist 2 entry 2 label",

		parsedPlaylists[1].Content.Items[1].Label,

		"should be able to parse playlist entry label",
	)

	assert.Equal(
		t,

		"playlist 3 entry 2 label",

		parsedPlaylists[2].Content.Items[1].Label,

		"should be able to parse playlist entry label",
	)

	os.RemoveAll("./playlists")
}
