package utils

import (
	"io/ioutil"
	"strings"
)

func ParseRetroarchPlaylistsFiles(files []FileInfo) ([]PlaylistInfo, error) {
	var playlists []PlaylistInfo = []PlaylistInfo{}

	for _, file := range files {
		fileStat := file.FsFile

		fileName := fileStat.Name()

		if !strings.HasSuffix(fileName, ".lpl") || fileStat.IsDir() {
			continue
		}

		plan, err := ioutil.ReadFile(file.Path)

		if err != nil {
			return nil, err
		}

		var playlistInfo = PlaylistInfo{
			Path: file.Path,

			Content: UnmarshalRetroarchPlaylist(plan),
		}

		playlists = AppendUniquePlaylist(playlists, playlistInfo)
	}

	return playlists, nil
}
