package utils

import (
	"errors"
	"path/filepath"
)

type LinkInfo struct {
	Name string

	Exe string

	Arguments string

	Content string
}

var ErrorInvalidCore = errors.New("alert, invalid core")

func GetLinkFileInfosFromPlaylistItem(
	playlistItem RetroArchPlaylistItem,

	retroArchExecutablePath string,

	outputLinksPath string,
) (LinkInfo, error) {
	if playlistItem.CorePath == "DETECT" {
		return LinkInfo{}, ErrorInvalidCore
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

	return infos, err
}
