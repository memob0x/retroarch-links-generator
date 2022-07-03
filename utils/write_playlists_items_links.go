package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

			outputPathStat, err := os.Stat(outputLinksPath)

			var outputPathDoesntExist bool = os.IsNotExist(err)

			var outputPathExists bool = !outputPathDoesntExist

			if outputPathDoesntExist && os.Mkdir(outputLinksPath, 0755) != nil {
				return errors.New("Error creating output folder.")
			}

			if outputPathExists && !outputPathStat.IsDir() {
				return errors.New("Output path is an existent file.")
			}

			linkPath, err := filepath.Abs(outputLinksPath + "/" + GetValidWinOsFilename(linkInfo.Name))

			// TODO: handle shortcut icon

			if runtime.GOOS == "windows" {
				err = CreateWinOsExeShortcut(
					retroArchExecutablePath,

					linkPath+".lnk",

					linkInfo.Arguments,
				)
			} else {
				err = ioutil.WriteFile(linkPath, []byte(linkInfo.Content), 0777)
			}

			if err != nil {
				return errors.New("The writing of file " + linkPath + " for playlist " + playlist.Path + " and rom " + playlistItem.RomPath + " returned the error: " + err.Error() + "\n")
			}
		}
	}

	return nil
}
