package utils

import (
	"encoding/json"
	"math"

	"github.com/wakeful-cloud/vdf"
)

func ConvertShortcutToVdfMap(shortcut SteamShortcut) vdf.Map {
	shortcutJson, _ := json.Marshal(&shortcut)

	var shortcutMapUnsafe vdf.Map = vdf.Map{}

	var shortcutMapSafe vdf.Map = vdf.Map{}

	_ = json.Unmarshal(shortcutJson, &shortcutMapUnsafe)

	for index, a := range shortcutMapUnsafe {
		switch a.(type) {
		case float64:
			shortcutMapSafe[index] = uint32(math.Round(a.(float64)))
		case string:
			shortcutMapSafe[index] = a
		}
	}

	return shortcutMapSafe
}
