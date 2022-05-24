package utils

import (
	"encoding/json"
	"io/ioutil"
)

type RetroArchPlaylistItem struct {
	RomPath string `json:"path"`

	Label string `json:"label"`

	CorePath string `json:"core_path"`
}

type RetroArchPlaylist struct {
	Items []RetroArchPlaylistItem `json:"items"`
}

func ParseRetroarchPlaylistsInFolder(playlistsPath string) ([]RetroArchPlaylist, error) {
	files, err := ioutil.ReadDir(playlistsPath)

	if err != nil {
		return nil, err
	}

	var playlists []RetroArchPlaylist = []RetroArchPlaylist{}

	for _, file := range files {
		if !file.IsDir() {
			plan, err := ioutil.ReadFile(playlistsPath + "/" + file.Name())

			if err != nil {
				return nil, err
			}

			var playlist RetroArchPlaylist

			json.Unmarshal(plan, &playlist)

			playlists = append(playlists, playlist)
		}
	}

	return playlists, nil
}
