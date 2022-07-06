package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/wakeful-cloud/vdf"
)

func WritePlaylistsItemsVdf(retroArchPlaylists []PlaylistInfo, retroArchExecutablePath string, vdfFilePath string) ([]SteamShortcut, error) {
	shortcuts, err := ReadShortcutsVdfFile(vdfFilePath)

	if err != nil && !errors.Is(err, ErrorVdfNotExists) {
		return nil, err
	}

	for _, playlist := range retroArchPlaylists {
		for _, playlistItem := range playlist.Content.Items {
			linkInfo, err := GetLinkFileInfosFromPlaylistItem(playlistItem, retroArchExecutablePath, "")

			if err != nil {
				fmt.Print("Playlist ", playlist.Path, " parsing for rom ", playlistItem.RomPath, " returned the error: ", err, "\n")

				continue
			}

			var shortcut SteamShortcut = SteamShortcut{
				AppId: GenerateSteamAppID(linkInfo.Exe + " " + linkInfo.Arguments),

				AppName: linkInfo.Name,

				// TODO: handle icon
				Icon: "",

				Exe: linkInfo.Exe,

				StartDir: linkInfo.Directory,

				LaunchOptions: linkInfo.Arguments,
			}

			if !SteamShortcutExists(shortcuts, shortcut) {
				shortcuts = append(shortcuts, shortcut)
			}
		}
	}

	shortcutsMap := ConvertShortcutsToVdfMap(shortcuts)

	for k, v := range shortcutsMap {
		fmt.Printf("%v %v", k, v)
	}

	vdfFileUpdated, err := vdf.WriteVdf(shortcutsMap)

	if err != nil {
		return nil, err
	}

	err = os.WriteFile(vdfFilePath, vdfFileUpdated, 0655)

	if err != nil {
		return nil, err
	}

	return shortcuts, nil
}
