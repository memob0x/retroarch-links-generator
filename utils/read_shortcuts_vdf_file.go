package utils

import (
	"encoding/json"
	"errors"

	"github.com/wakeful-cloud/vdf"
)

type SteamShortcut struct {
	AppId         uint32 `json:"appid"`
	AppName       string `json:"AppName"`
	Exe           string `json:"Exe"`
	LaunchOptions string `json:"LaunchOptions"`
	Icon          string `json:"icon"`
	StartDir      string `json:"StartDir"`

	// unused

	AllowOverlay        uint32  `json:"AllowOverlay"`
	OpenVR              uint32  `json:"OpenVR"`
	FlatpakAppID        string  `json:"FlatpakAppID"`
	IsHidden            uint32  `json:"IsHidden"`
	AllowDesktopConfig  uint32  `json:"AllowDesktopConfig"`
	DevkitGameID        string  `json:"DevkitGameID"`
	LastPlayTime        uint32  `json:"LastPlayTime"`
	ShortcutPath        string  `json:"ShortcutPath"`
	DevkitOverrideAppID uint32  `json:"DevkitOverrideAppID"`
	Devkit              uint32  `json:"Devkit"`
	Tags                vdf.Map `json:"tags"`
}

func ReadShortcutsVdfFile(vdfFilePath string) ([]SteamShortcut, error) {
	vdfMap, err := ReadBinaryVdfFile(vdfFilePath)

	if err != nil {
		return nil, err
	}

	shortcutsMap, isCastOk := vdfMap["shortcuts"].(vdf.Map)

	if !isCastOk {
		return nil, errors.New("Not a shortcuts vdf")
	}

	var shortcuts []SteamShortcut

	for _, shortcutMapEntry := range shortcutsMap {
		shortcutMapValue, _ := shortcutMapEntry.(vdf.Map)

		var shortcut SteamShortcut

		shortcutMapValueBytes, _ := json.Marshal(shortcutMapValue)

		_ = json.Unmarshal(shortcutMapValueBytes, &shortcut)

		shortcuts = append(shortcuts, shortcut)
	}

	return shortcuts, ErrorVdfNotExists
}
