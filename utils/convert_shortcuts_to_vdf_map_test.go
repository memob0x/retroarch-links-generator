package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wakeful-cloud/vdf"
)

func TestConvertShortcutsToVdfMap(t *testing.T) {
	shortcutVdfMap, _ := ConvertShortcutsToVdfMap([]SteamShortcut{
		{
			AppId:         123,
			AppName:       "1",
			Exe:           "2",
			LaunchOptions: "3",
			Icon:          "4",
			StartDir:      "5",
		},
		{
			AppId:         345,
			AppName:       "1",
			Exe:           "2",
			LaunchOptions: "3",
			Icon:          "4",
			StartDir:      "5",
		},
	})

	assert.Equal(
		t,

		vdf.Map{
			"shortcuts": vdf.Map{
				"0": vdf.Map{
					"appid":               uint32(123),
					"AppName":             "1",
					"Exe":                 "2",
					"LaunchOptions":       "3",
					"StartDir":            "5",
					"icon":                "4",
					"ShortcutPath":        "",
					"FlatpakAppID":        "",
					"DevkitGameID":        "",
					"AllowDesktopConfig":  uint32(0),
					"AllowOverlay":        uint32(0),
					"Devkit":              uint32(0),
					"DevkitOverrideAppID": uint32(0),
					"IsHidden":            uint32(0),
					"LastPlayTime":        uint32(0),
					"OpenVR":              uint32(0),
				},
				"1": vdf.Map{
					"appid":               uint32(345),
					"AppName":             "1",
					"Exe":                 "2",
					"LaunchOptions":       "3",
					"StartDir":            "5",
					"icon":                "4",
					"ShortcutPath":        "",
					"FlatpakAppID":        "",
					"DevkitGameID":        "",
					"AllowDesktopConfig":  uint32(0),
					"AllowOverlay":        uint32(0),
					"Devkit":              uint32(0),
					"DevkitOverrideAppID": uint32(0),
					"IsHidden":            uint32(0),
					"LastPlayTime":        uint32(0),
					"OpenVR":              uint32(0),
				},
			},
		},

		shortcutVdfMap,

		"Should be able to convert many steam shortcuts to a vdf maps",
	)
}
