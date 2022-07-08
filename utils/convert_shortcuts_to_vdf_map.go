package utils

import (
	"strconv"

	"github.com/wakeful-cloud/vdf"
)

func ConvertShortcutsToVdfMap(shortcuts []SteamShortcut) (vdf.Map, error) {
	var shortcutsMap vdf.Map = vdf.Map{}

	for index, shortcut := range shortcuts {
		shortcutMap, err := ConvertShortcutToVdfMap(shortcut)

		if err != nil {
			return nil, err
		}

		shortcutsMap[strconv.Itoa(index)] = shortcutMap
	}

	return vdf.Map{
		"shortcuts": shortcutsMap,
	}, nil
}
