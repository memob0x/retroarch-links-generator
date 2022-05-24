package utils

import "errors"

type LinkInfo struct {
	Content string

	Path string
}

func GetLinkFileInfosFromPlaylistItem(playlistItem RetroArchPlaylistItem, src string, dist string) (LinkInfo, error) {
	if playlistItem.CorePath == "DETECT" {
		return LinkInfo{}, errors.New("Alert, invalid core " + playlistItem.CorePath + " for " + playlistItem.RomPath)
	}

	var infos LinkInfo

	var executablePath string = src + "\\retroarch.exe"

	infos.Content = executablePath +
		" " +
		"--load-menu-on-error" +
		" " +
		"-L" +
		" " +
		playlistItem.CorePath +
		" " +
		playlistItem.RomPath

	infos.Path = dist + "\\" + GetValidWinOsFilename(playlistItem.Label) + ".bat"

	return infos, nil
}
