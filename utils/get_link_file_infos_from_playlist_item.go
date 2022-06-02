package utils

import (
	"errors"
	"path/filepath"
	"runtime"
)

type LinkInfo struct {
	Content string

	Path string
}

func GetLinkFileInfosFromPlaylistItem(
	playlistItem RetroArchPlaylistItem,

	retroArchExecutablePath string,

	outputLinksPath string,
) (LinkInfo, error) {
	var linkExtension = ""

	if runtime.GOOS == "windows" {
		linkExtension = ".bat"
	}

	if playlistItem.CorePath == "DETECT" {
		return LinkInfo{}, errors.New("Alert, invalid core " + playlistItem.CorePath + " for " + playlistItem.RomPath)
	}

	var infos LinkInfo

	executablePath, err := filepath.Abs(retroArchExecutablePath)

	infos.Content = executablePath +
		" " +
		"--load-menu-on-error" +
		" " +
		"-L" +
		" " +
		"\"" + playlistItem.CorePath + "\"" +
		" " +
		"\"" + playlistItem.RomPath + "\""

	linkPath, err := filepath.Abs(outputLinksPath + "/" + GetValidWinOsFilename(playlistItem.Label) + linkExtension)

	infos.Path = linkPath

	return infos, err
}
