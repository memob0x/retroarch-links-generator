package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var ErrorLinkCreation error = errors.New("Link creation error")
var ErrorOutputFolder error = errors.New("Error creating output folder.")
var ErrorExistentFile error = errors.New("Output path is an existent file.")

func WritePlaylistsItemsLinks(playlists []PlaylistInfo, retroArchExecutablePath string, outputLinksPath string) error {
	for _, playlist := range playlists {
		for _, playlistItem := range playlist.Content.Items {
			linkInfo, err := GetLinkFileInfosFromPlaylistItem(playlistItem, retroArchExecutablePath, outputLinksPath)

			var warningMessage = fmt.Sprintf(
				"Playlist %v parsing for rom %v returned the error: %v\n",

				playlist.Path,

				playlistItem.RomPath,

				err,
			)

			if err != nil {
				fmt.Print(warningMessage)

				continue
			}

			outputPathStat, err := os.Stat(outputLinksPath)

			var outputPathDoesntExist bool = os.IsNotExist(err)

			var outputPathExists bool = !outputPathDoesntExist

			if outputPathDoesntExist && os.Mkdir(outputLinksPath, 0755) != nil {
				fmt.Print(warningMessage)

				return ErrorOutputFolder
			}

			if outputPathExists && !outputPathStat.IsDir() {
				fmt.Print(warningMessage)

				return ErrorExistentFile
			}

			linkPath, err := filepath.Abs(outputLinksPath + "/" + GetValidWinOsFilename(linkInfo.Name))

			if err != nil {
				return ErrorLinkCreation
			}

			// TODO: maybe handle icon

			if IS_OS_WINDOWS {
				err = CreateWinOsExeShortcut(
					retroArchExecutablePath,

					linkPath+".lnk",

					linkInfo.Arguments,
				)
			} else {
				err = ioutil.WriteFile(linkPath, []byte(linkInfo.Content), 0777)
			}

			if err != nil {
				return ErrorLinkCreation
			}
		}
	}

	return nil
}
