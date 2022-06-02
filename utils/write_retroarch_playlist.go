package utils

import (
	"encoding/json"
	"os"
)

func WriteRetroarchPlaylist(playlist RetroArchPlaylist, path string) (*os.File, int, error) {
	f, err := os.Create(path)

	if err != nil {
		return f, 0, err
	}

	defer f.Close()

	playlistMarshaled, _ := json.Marshal(playlist)

	n, err := f.Write(playlistMarshaled)

	if err != nil {
		return f, 0, err
	}

	return f, n, err
}
