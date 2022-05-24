package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

type RetroArchPlaylistItem struct {
	RomPath string `json:"path"`

	Label string `json:"label"`

	CorePath string `json:"core_path"`
}

type RetroArchPlaylist struct {
	Items []RetroArchPlaylistItem `json:"items"`
}

// thanks https://stackoverflow.com/questions/32438204/create-a-windows-shortcut-lnk-in-go
func MakeWinOSLink(src string, dst string) error {
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)

	oleShellObject, err := oleutil.CreateObject("WScript.Shell")

	if err != nil {
		return err
	}

	defer oleShellObject.Release()

	wshell, err := oleShellObject.QueryInterface(ole.IID_IDispatch)

	if err != nil {
		return err
	}

	defer wshell.Release()

	cs, err := oleutil.CallMethod(wshell, "CreateShortcut", dst)

	if err != nil {
		return err
	}

	idispatch := cs.ToIDispatch()

	oleutil.PutProperty(idispatch, "TargetPath", src)

	oleutil.CallMethod(idispatch, "Save")

	return nil
}

func ParseRetroArchPlaylists(playlistsPath string) []RetroArchPlaylist {
	files, err := ioutil.ReadDir(playlistsPath)

	if err != nil {
		log.Fatal(err)
	}

	var playlists []RetroArchPlaylist = []RetroArchPlaylist{}

	for _, file := range files {
		if !file.IsDir() {
			plan, err := ioutil.ReadFile(playlistsPath + "/" + file.Name())

			if err != nil {
				log.Fatal(err)
			}

			var playlist RetroArchPlaylist

			json.Unmarshal(plan, &playlist)

			playlists = append(playlists, playlist)
		}
	}

	return playlists
}

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

	for _, playlist := range ParseRetroArchPlaylists(path.Dir(retroArchExecutablePath) + "/playlists") {
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

				MakeWinOSLink(linkCmd, outputLinksPath)
			}
		}
	}
}
