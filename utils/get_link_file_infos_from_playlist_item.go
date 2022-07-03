package utils

import (
	"errors"
	"path/filepath"
)

type LinkInfo struct {
	Name string

	Exe string

	Directory string

	Arguments string

	Content string
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

	infos.Exe = retroArchExecutablePath

	infos.Content = executablePath +
		" " +
		infos.Arguments

	infos.Name = playlistItem.Label

	infos.Directory = outputLinksPath

	return infos, err
}
