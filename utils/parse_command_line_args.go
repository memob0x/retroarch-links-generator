package utils

import (
	"errors"
	"path/filepath"
)

func ParseCommandLineArgs(args []string) (string, string, string, error) {
	var argsCount = len(args)

	var retroArchExecutablePath string = ""

	linksOutputDestPath, err := filepath.Abs(".")

	if argsCount >= 2 {
		retroArchExecutablePath, err = filepath.Abs(args[1])
	} else {
		err = errors.New("No retroarch executable path specified.")
	}

	retroArchPlaylistsPath, err := filepath.Abs(filepath.Dir(retroArchExecutablePath) + "/playlists")

	if argsCount >= 3 {
		linksOutputDestPath, err = filepath.Abs(args[2])
	}

	if argsCount >= 4 {
		retroArchPlaylistsPath, err = filepath.Abs(args[3])
	}

	return retroArchExecutablePath, linksOutputDestPath, retroArchPlaylistsPath, err
}
