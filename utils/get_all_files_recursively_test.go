package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllFilesRecursively(t *testing.T) {
	os.Mkdir("./playlists", 0755)
	os.Mkdir("./playlists/foo", 0755)
	os.Mkdir("./playlists/foo/bar", 0755)
	os.Mkdir("./playlists/foo/bar/biz/", 0755)

	WriteRetroarchPlaylist(RetroArchPlaylist{}, "./playlists/playlist.lpl")
	WriteRetroarchPlaylist(RetroArchPlaylist{}, "./playlists/foo/playlist.lpl")
	WriteRetroarchPlaylist(RetroArchPlaylist{}, "./playlists/foo/bar/biz/playlist.lpl")

	playlistsDirStat, _ := os.Stat("./playlists")

	e, _ := GetAllFilesRecursively([]FileInfo{
		{
			Path: "./playlists",

			FsFile: playlistsDirStat,
		},
	})

	assert.Equal(
		t,

		len(e),

		3,

		"Should be able to collect all files (not folders) recursively",
	)

	playlistFileStat, _ := os.Stat("./playlists/playlist.lpl")

	f, _ := GetAllFilesRecursively([]FileInfo{
		{
			Path: "./playlists/playlist.lpl",

			FsFile: playlistFileStat,
		},
	})

	assert.Equal(
		t,

		len(f),

		1,

		"Should be able to collect all files (not folders) recursively",
	)

	os.Remove("./playlists")
}
