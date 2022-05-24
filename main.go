package main

import (
	"fmt"
	"log"
	"os"
	"path"

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
		fmt.Printf("No output folder specified, fallback to \"%v\".\n", linksDestPath)
	}

	return retroArchPath, linksDestPath
}

func main() {
	var retroArchExecutablePath, outputLinksPath = ParseCommandLineArgs()

	for _, playlist := range utils.ParseRetroarchPlaylistsInFolder(path.Dir(retroArchExecutablePath) + "/playlists") {
		for _, playlistItem := range playlist.Items {
			if playlistItem.CorePath != "DETECT" {
				var linkCmd string = retroArchExecutablePath +
					" " +
					"--load-menu-on-error" +
					" " +
					"-L" +
					" " +
					playlistItem.CorePath +
					" " +
					playlistItem.RomPath

				fmt.Print(linkCmd, "\n")

				utils.MakeWinOSLink(linkCmd, outputLinksPath)
			}
		}
	}
}
