package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wakeful-cloud/vdf"
)

func TestWritePlaylistsItemsVdf(t *testing.T) {
	WritePlaylistsItemsVdf(
		[]PlaylistInfo{
			{
				Path: "/fake/path",

				Content: RetroArchPlaylist{
					Items: []RetroArchPlaylistItem{
						{
							RomPath: "a",

							Label: "b",

							CorePath: "c",
						},
					},
				},
			},
		},

		"../data/executable",

		"./written-playlists.vdf",
	)

	file, _ := os.ReadFile("./written-playlists.vdf")

	vdfMap, _ := vdf.ReadVdf(file)

	pathAbs, _ := filepath.Abs("../data/")

	assert.Equal(
		t,

		vdf.Map{
			"shortcuts": vdf.Map{
				"0": vdf.Map{
					"AppName":             "b",
					"LaunchOptions":       "--load-menu-on-error -L \"c\" \"a\"",
					"Exe":                 "../data/executable",
					"StartDir":            pathAbs + "/",
					"appid":               GenerateSteamAppID("../data/executable --load-menu-on-error -L \"c\" \"a\""),
					"icon":                "",
					"DevkitGameID":        "",
					"FlatpakAppID":        "",
					"ShortcutPath":        "",
					"IsHidden":            uint32(0),
					"DevkitOverrideAppID": uint32(0),
					"Devkit":              uint32(0),
					"LastPlayTime":        uint32(0),
					"AllowDesktopConfig":  uint32(0),
					"AllowOverlay":        uint32(0),
					"OpenVR":              uint32(0),
				},
			},
		},

		vdfMap,

		"Should be able to write RetroArch playlists as Steam shortcuts vdf file",
	)

	os.Remove("./written-playlists.vdf")
}
