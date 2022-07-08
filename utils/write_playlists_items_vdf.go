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

			var warningMessage string = fmt.Sprintf("Playlist %v parsing for rom %v returned the error: %v\n", playlist.Path, playlistItem.RomPath, err)

			if err != nil {
				fmt.Print(warningMessage)

				continue
			}

			startDir, err := GetDirPath(retroArchExecutablePath)

			if err != nil {
				fmt.Print(warningMessage)

				continue
			}

			var shortcut SteamShortcut = SteamShortcut{
				AppId: GenerateSteamAppID(linkInfo.Exe + " " + linkInfo.Arguments),

				AppName: linkInfo.Name,

				// TODO: maybe handle icon
				Icon: "",

				Exe: linkInfo.Exe,

				StartDir: startDir,

				LaunchOptions: linkInfo.Arguments,
			}

			if !SteamShortcutExists(shortcuts, shortcut) {
				shortcuts = append(shortcuts, shortcut)
			}
		}
	}

	shortcutsMap, err := ConvertShortcutsToVdfMap(shortcuts)

	if err != nil {
		return nil, err
	}

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
