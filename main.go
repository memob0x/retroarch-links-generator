package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/memob0x/retroarch-links-generator/utils"
)

func main() {
	var retroArchPath, outputLinksPath = utils.ParseCommandLineArgs(os.Args)

	playlists, err := utils.ParseRetroarchPlaylistsInPath(retroArchPath)

	if err != nil {
		log.Fatal(err)
	}

	for _, playlist := range playlists {
		for _, playlistItem := range playlist.Content.Items {
			linkInfo, err := utils.GetLinkFileInfosFromPlaylistItem(playlistItem, retroArchPath, outputLinksPath)

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
