package utils

import (
	"os"
	"strings"
)

type PlaylistInfo struct {
	Path string

	Content RetroArchPlaylist
}

func ParseRetroarchPlaylistsPaths(playlistsValue string) ([]PlaylistInfo, error) {
	playlistsValues := strings.Split(playlistsValue, ",")

	var allPlaylists []PlaylistInfo = []PlaylistInfo{}

	for _, playlistValue := range playlistsValues {
		stat, err := os.Stat(playlistValue)

		if err != nil {
			return nil, err
		}

		files, err := GetAllFilesRecursively([]FileInfo{
			{
				Path: playlistValue,

				FsFile: stat,
			},
		})

		if err != nil {
			return nil, err
		}

		filesPlaylists, err := ParseRetroarchPlaylistsFiles(files)

		if err != nil {
			return nil, err
		}

		allPlaylists = AppendUniquePlaylists(allPlaylists, filesPlaylists)
	}

	return allPlaylists, nil
}
