package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/memob0x/retroarch-links-generator/utils"
)

func ParseCommandLineArgs() (string, string) {
	var argsCount = len(os.Args)

	var retroArchPath string

	var linksDestPath string = "."

	if argsCount >= 2 {
		retroArchPath = os.Args[1]
	} else {
		log.Fatalf("No retroarch path specified, aborting.\n")
	}

	if argsCount >= 3 {
		linksDestPath = os.Args[2]
	} else {
		log.Fatalf("No destination path specified, aborting.\n")
	}

	return retroArchPath, linksDestPath
}

func main() {
	var retroArchPath, outputLinksPath = ParseCommandLineArgs()

	playlists, err := utils.ParseRetroarchPlaylistsInFolder(retroArchPath + "/playlists")

	if err != nil {
		log.Fatal(err)
	}

	var executablePath string = retroArchPath + "\\retroarch.exe"

	for _, playlist := range playlists {
		for _, playlistItem := range playlist.Items {
			if playlistItem.CorePath != "DETECT" {
				var cmd string = executablePath +
					" " +
					"--load-menu-on-error" +
					" " +
					"-L" +
					" " +
					playlistItem.CorePath +
					" " +
					playlistItem.RomPath

				err := ioutil.WriteFile(outputLinksPath+"\\"+playlistItem.Label+".bat", []byte(cmd), 0644)

				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
