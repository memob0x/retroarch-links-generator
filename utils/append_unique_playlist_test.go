package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendUniquePlaylist(t *testing.T) {
	playlists := []PlaylistInfo{
		{
			Path: "a",

			Content: RetroArchPlaylist{
				Items: []RetroArchPlaylistItem{},
			},
		},
		{
			Path: "b",

			Content: RetroArchPlaylist{
				Items: []RetroArchPlaylistItem{},
			},
		},
	}

	a := AppendUniquePlaylist(
		playlists,

		PlaylistInfo{
			Path: "b",

			Content: RetroArchPlaylist{
				Items: []RetroArchPlaylistItem{},
			},
		},
	)

	assert.Equal(
		t,

		a,

		playlists,

		"Should be able avoid duplicates",
	)

	c := AppendUniquePlaylist(
		playlists,

		PlaylistInfo{
			Path: "c",

			Content: RetroArchPlaylist{
				Items: []RetroArchPlaylistItem{},
			},
		},
	)

	assert.Equal(
		t,

		c,

		[]PlaylistInfo{
			{
				Path: "a",

				Content: RetroArchPlaylist{
					Items: []RetroArchPlaylistItem{},
				},
			},
			{
				Path: "b",

				Content: RetroArchPlaylist{
					Items: []RetroArchPlaylistItem{},
				},
			},
			{
				Path: "c",

				Content: RetroArchPlaylist{
					Items: []RetroArchPlaylistItem{},
				},
			},
		},

		"Should be able avoid duplicates",
	)
}
