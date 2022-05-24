package utils

import (
	"io/ioutil"
)

type PlaylistInfo struct {
	Path string

	Content RetroArchPlaylist
}

func ParseRetroarchPlaylistsInPath(path string) ([]PlaylistInfo, error) {
	var playlistsPath string = path + "/playlists"

	files, err := ioutil.ReadDir(playlistsPath)

	if err != nil {
		return nil, err
	}

	var playlists []PlaylistInfo = []PlaylistInfo{}

	for _, file := range files {
		if !file.IsDir() {
			var playlistPath string = playlistsPath + "/" + file.Name()

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
