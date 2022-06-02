package utils

import (
	"io/ioutil"
	"path/filepath"
)

type PlaylistInfo struct {
	Path string

	Content RetroArchPlaylist
}

func ParseRetroarchPlaylistsInPath(playlistsPath string) ([]PlaylistInfo, error) {
	files, err := ioutil.ReadDir(playlistsPath)

	if err != nil {
		return nil, err
	}

	var playlists []PlaylistInfo = []PlaylistInfo{}

	for _, file := range files {
		if !file.IsDir() {
			playlistPath, err := filepath.Abs(playlistsPath + "/" + file.Name())

			if err != nil {
				return nil, err
			}

			plan, err := ioutil.ReadFile(playlistPath)

			if err != nil {
				return nil, err
			}

			var playlistInfo = PlaylistInfo{
				Path: playlistPath,

				Content: UnmarshalRetroarchPlaylist(plan),
			}

			playlists = append(playlists, playlistInfo)
		}
	}

	return playlists, nil
}
