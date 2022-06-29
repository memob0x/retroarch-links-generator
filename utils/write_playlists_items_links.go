package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"runtime"
)

func WritePlaylistsItemsLinks(playlists []PlaylistInfo, retroArchExecutablePath string, outputLinksPath string) error {
	for _, playlist := range playlists {
		for _, playlistItem := range playlist.Content.Items {
			linkInfo, err := GetLinkFileInfosFromPlaylistItem(playlistItem, retroArchExecutablePath, outputLinksPath)

			if err != nil {
				fmt.Print("Playlist ", playlist.Path, " parsing for rom ", playlistItem.RomPath, " returned the error: ", err, "\n")

				continue
			}

			if runtime.GOOS == "windows" {
				err = CreateWinOsExeShortcut(
					retroArchExecutablePath,

					linkInfo.Path+".lnk",

					linkInfo.Arguments,
				)
			} else {
				err = ioutil.WriteFile(linkInfo.Path, []byte(linkInfo.Content), 0777)
			}

			if err != nil {
				return errors.New("The writing of file " + linkInfo.Path + " for playlist " + playlist.Path + " and rom " + playlistItem.RomPath + " returned the error: " + err.Error() + "\n")
			}
		}
	}

	return nil
}
