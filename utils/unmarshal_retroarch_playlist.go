package utils

import (
	"encoding/json"
)

type RetroArchPlaylistItem struct {
	RomPath string `json:"path"`

	Label string `json:"label"`

	CorePath string `json:"core_path"`
}

type RetroArchPlaylist struct {
	Items []RetroArchPlaylistItem `json:"items"`
}

func UnmarshalRetroarchPlaylist(value []byte) RetroArchPlaylist {
	var playlist RetroArchPlaylist

	json.Unmarshal(value, &playlist)

	return playlist
}
