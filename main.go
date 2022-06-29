package main

import (
	"log"
	"os"

	"github.com/memob0x/retroarch-links-generator/utils"
)

func MaybePanic(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var retroArchExecutablePath, outputPath, retroArchPlaylistsFolderPath, isOutputPathVdfFile, err = utils.ParseCommandLineArgs(os.Args)

	MaybePanic(err)

	playlists, err := utils.ParseRetroarchPlaylistsInPath(retroArchPlaylistsFolderPath)

	MaybePanic(err)

	if isOutputPathVdfFile {
		err = utils.WritePlaylistsItemsVdf(playlists, retroArchExecutablePath, outputPath)

		MaybePanic(err)

		return
	}

	if !isOutputPathVdfFile {
		err = utils.WritePlaylistsItemsLinks(playlists, retroArchExecutablePath, outputPath)

		MaybePanic(err)

		return
	}
}
