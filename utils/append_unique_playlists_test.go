package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendUniquePlaylists(t *testing.T) {
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

	a := AppendUniquePlaylists(
		playlists,

		playlists,
	)

	assert.Equal(
		t,

		a,

		playlists,

		"Should be able avoid duplicates",
	)

	playlists2 := []PlaylistInfo{
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
	}

	c := AppendUniquePlaylists(
		playlists,

		playlists2,
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
