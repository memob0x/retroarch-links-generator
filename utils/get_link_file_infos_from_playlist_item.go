package utils

import (
	"errors"
	"path/filepath"
)

type LinkInfo struct {
	Directory string

	Arguments string

	Content string

	Path string
}

func GetLinkFileInfosFromPlaylistItem(
	playlistItem RetroArchPlaylistItem,

	retroArchExecutablePath string,

	outputLinksPath string,
) (LinkInfo, error) {
	if playlistItem.CorePath == "DETECT" {
		return LinkInfo{}, errors.New("Alert, invalid core " + playlistItem.CorePath + " for " + playlistItem.RomPath)
	}

	var infos LinkInfo

	executablePath, err := filepath.Abs(retroArchExecutablePath)

	infos.Arguments = "--load-menu-on-error" +
		" " +
		"-L" +
		" " +
		"\"" + playlistItem.CorePath + "\"" +
		" " +
		"\"" + playlistItem.RomPath + "\""

	infos.Content = executablePath +
		" " +
		infos.Arguments

	linkPath, err := filepath.Abs(outputLinksPath + "/" + GetValidWinOsFilename(playlistItem.Label))

	infos.Directory = outputLinksPath

	infos.Path = linkPath

	return infos, err
}
