package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRetroarchPlaylistsInPath(t *testing.T) {
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

	parsedPlaylists, _ := ParseRetroarchPlaylistsInPath("./")

	assert.Equal(
		t,

		3,

		len(parsedPlaylists),

		"should be able to parse all playlists in \"playlists\" in folder",
	)

	assert.Equal(
		t,

		".//playlists/playlist-1.lpl", // FIXME: fix double slash

		parsedPlaylists[0].Path,

		"should be able to return the parsed playlist path",
	)

	assert.Equal(
		t,

		2,

		len(parsedPlaylists[0].Content.Items),

		"should be able to parse all the given playlist entries",
	)

	assert.Equal(
		t,

		"/path/to/core",

		parsedPlaylists[0].Content.Items[0].CorePath,

		"should be able to parse playlist core path",
	)

	assert.Equal(
		t,

		"/path/to/rom",

		parsedPlaylists[0].Content.Items[0].RomPath,

		"should be able to parse playlist rom path",
	)

	assert.Equal(
		t,

		"playlist 1 entry 1 label",

		parsedPlaylists[0].Content.Items[0].Label,

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
