package utils

import (
	"errors"
	"path/filepath"
	"strings"
)

func ParseCommandLineArgs(args []string) (string, string, string, bool, error) {
	var argsCount = len(args)

	var retroArchExecutablePath string = ""

	outputDestPath, err := filepath.Abs(".")

	if argsCount >= 2 {
		retroArchExecutablePath, err = filepath.Abs(args[1])
	} else {
		err = errors.New("No retroarch executable path specified.")
	}

	retroArchPlaylists, err := filepath.Abs(filepath.Dir(retroArchExecutablePath) + "/playlists")

	if argsCount >= 3 {
		outputDestPath, err = filepath.Abs(args[2])
	}

	var isOutputPathVdfFile bool = strings.HasSuffix(outputDestPath, ".vdf")

	if argsCount >= 4 {
		retroArchPlaylists, err = filepath.Abs(args[3])
	}

	return retroArchExecutablePath, outputDestPath, retroArchPlaylists, isOutputPathVdfFile, err
}
