package utils

func SteamShortcutExists(shortcuts []SteamShortcut, shortcut SteamShortcut) bool {
	for _, existentShortcut := range shortcuts {
		if existentShortcut.AppName == shortcut.AppName || existentShortcut.AppId == shortcut.AppId {
			return true
		}
	}

	return false
}
