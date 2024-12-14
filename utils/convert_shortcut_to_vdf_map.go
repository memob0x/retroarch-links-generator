package utils

import (
	"encoding/json"
	"errors"
	"math"

	"github.com/wakeful-cloud/vdf"
)

var ErrorInvalidVdfProp error = errors.New("invalid vdf type property")

func ConvertShortcutToVdfMap(shortcut SteamShortcut) (vdf.Map, error) {
	shortcutJson, _ := json.Marshal(&shortcut)

	var shortcutMapUnsafe vdf.Map = vdf.Map{}

	var shortcutMapSafe vdf.Map = vdf.Map{}

	_ = json.Unmarshal(shortcutJson, &shortcutMapUnsafe)

	for index, a := range shortcutMapUnsafe {
		// FIXME: "tags" (vdf.Map) ends in switch "default" case somehow
		if index == "tags" {
			continue
		}

		switch b := a.(type) {
		case float64:
			shortcutMapSafe[index] = uint32(math.Round(b))
		case uint32, string, vdf.Map:
			shortcutMapSafe[index] = a
		default:
			return nil, ErrorInvalidVdfProp
		}
	}

	return shortcutMapSafe, nil
}
