package utils

import (
	"os"
	"path/filepath"
	"regexp"
)

var FilenameRegexp *regexp.Regexp = regexp.MustCompile("[A-Za-z0-9-_.]+$")

func GetDirPath(inputPath string) (string, error) {
	pathAbsolute, err := filepath.Abs(inputPath)

	if err != nil {
		return "", err
	}

	fileInfo, err := os.Stat(pathAbsolute)

	if err != nil {
		return "", err
	}

	if fileInfo.IsDir() {
		return pathAbsolute, nil
	}

	return FilenameRegexp.ReplaceAllString(pathAbsolute, ""), nil
}
