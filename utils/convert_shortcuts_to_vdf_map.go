package utils

import (
	"strconv"

	"github.com/wakeful-cloud/vdf"
)

func ConvertShortcutsToVdfMap(shortcuts []SteamShortcut) vdf.Map {
	var shortcutsMap vdf.Map = vdf.Map{}

	for index, shortcut := range shortcuts {
		shortcutsMap[strconv.Itoa(index)] = ConvertShortcutToVdfMap(shortcut)
	}

	return vdf.Map{
		"shortcuts": shortcutsMap,
	}
}
