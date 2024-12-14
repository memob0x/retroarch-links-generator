package utils

import (
	"errors"
	"path/filepath"
	"strings"
)

var ErrorRetroArchExecutablePath = errors.New("no retroarch executable path specified")

func ParseCommandLineArgs(args []string) (string, string, string, bool, error) {
	var argsCount = len(args)

	var retroArchExecutablePath string = ""

	outputDestPath, err := filepath.Abs(".")

	if err != nil {
		return "", "", "", false, err
	}

	if argsCount >= 2 {
		retroArchExecutablePath, err = filepath.Abs(args[1])
	} else {
		err = ErrorRetroArchExecutablePath
	}

	if err != nil {
		return "", "", "", false, err
	}

	retroArchPlaylists, err := filepath.Abs(filepath.Dir(retroArchExecutablePath) + "/playlists")

	if err != nil {
		return "", "", "", false, err
	}

	if argsCount >= 3 {
		outputDestPath, err = filepath.Abs(args[2])
	}

	if err != nil {
		return "", "", "", false, err
	}

	var isOutputPathVdfFile bool = strings.HasSuffix(outputDestPath, ".vdf")

	if argsCount >= 4 {
		retroArchPlaylists, err = filepath.Abs(args[3])
	}

	return retroArchExecutablePath, outputDestPath, retroArchPlaylists, isOutputPathVdfFile, err
}
