package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/memob0x/retroarch-links-generator/utils"
)

func main() {
	var retroArchExecutablePath, outputLinksPath, retroArchPlaylistsFolderPath, err = utils.ParseCommandLineArgs(os.Args)

	if err != nil {
		log.Fatal(err)
	}

	playlists, err := utils.ParseRetroarchPlaylistsInPath(retroArchPlaylistsFolderPath)

	if err != nil {
		log.Fatal(err)
	}

	for _, playlist := range playlists {
		for _, playlistItem := range playlist.Content.Items {
			linkInfo, err := utils.GetLinkFileInfosFromPlaylistItem(playlistItem, retroArchExecutablePath, outputLinksPath)

			if err != nil {
				fmt.Print("Playlist ", playlist.Path, " parsing for rom ", playlistItem.RomPath, " returned the error: ", err, "\n")

				continue
			}

			err = ioutil.WriteFile(linkInfo.Path, []byte(linkInfo.Content), 0644)

			if err != nil {
				fmt.Print("The writing of file ", linkInfo.Path, " for playlist ", playlist.Path, " and rom ", playlistItem.RomPath, " returned the error: ", err, "\n")
			}
		}
	}
}
