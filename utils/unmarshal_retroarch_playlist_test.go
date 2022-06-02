package utils

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalRetroarchPlaylist(t *testing.T) {
	var playlist = RetroArchPlaylist{
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
	}

	playlistEncoded, _ := json.Marshal(playlist)

	var playlistDecoded = UnmarshalRetroarchPlaylist(playlistEncoded)

	assert.Equal(
		t,

		playlistDecoded,

		playlist,

		"should be able to unmarshal a given retroarch json encoded playlist",
	)
}
