package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRetroarchPlaylistsPaths(t *testing.T) {
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

	parsedPlaylists, _ := ParseRetroarchPlaylistsPaths("./playlists")

	assert.Equal(
		t,

		len(parsedPlaylists),

		3,

		"should be able to parse all playlists in \"playlists\" folder",
	)

	assert.Contains(
		t,

		parsedPlaylists[0].Path,

		"/playlists/playlist-1.lpl",

		"should be able to return the parsed playlist path",
	)

	assert.Equal(
		t,

		len(parsedPlaylists[0].Content.Items),

		2,

		"should be able to parse all the given playlist entries",
	)

	assert.Contains(
		t,

		parsedPlaylists[0].Content.Items[0].CorePath,

		"/path/to/core",

		"should be able to parse playlist core path",
	)

	assert.Contains(
		t,

		parsedPlaylists[0].Content.Items[0].RomPath,

		"/path/to/rom",

		"should be able to parse playlist rom path",
	)

	assert.Equal(
		t,

		parsedPlaylists[0].Content.Items[0].Label,

		"playlist 1 entry 1 label",

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

	parsedPlaylistsParts, _ := ParseRetroarchPlaylistsPaths("./playlists/playlist-1.lpl,./playlists/playlist-3.lpl")

	assert.Equal(
		t,

		len(parsedPlaylistsParts),

		2,

		"should be able to parse all playlists in \"playlists\" folder",
	)

	assert.Contains(
		t,

		parsedPlaylistsParts[0].Path,

		"/playlists/playlist-1.lpl",

		"should be able to return the parsed playlist path",
	)

	assert.Equal(
		t,

		len(parsedPlaylistsParts[0].Content.Items),

		2,

		"should be able to parse all the given playlist entries",
	)

	assert.Contains(
		t,

		parsedPlaylistsParts[0].Content.Items[0].CorePath,

		"/path/to/core",

		"should be able to parse playlist core path",
	)

	assert.Contains(
		t,

		parsedPlaylistsParts[0].Content.Items[0].RomPath,

		"/path/to/rom",

		"should be able to parse playlist rom path",
	)

	assert.Equal(
		t,

		parsedPlaylistsParts[0].Content.Items[0].Label,

		"playlist 1 entry 1 label",

		"should be able to parse playlist entry label",
	)

	assert.Equal(
		t,

		"playlist 3 entry 2 label",

		parsedPlaylistsParts[1].Content.Items[1].Label,

		"should be able to parse playlist entry label",
	)

	os.RemoveAll("./playlists")
}
