package utils

import (
	"errors"
	"fmt"
	"strings"
)

var ErrWinOsExeShortcutCreationInput = errors.New("invalid executable paths provided")

func GetVbsScriptShortcutCreatorContent(
	shortcutSourceFilePath string,

	shortcutTargetFilePath string,

	shortcutArguments string,
) (string, error) {
	if !strings.HasSuffix(shortcutSourceFilePath, ".exe") || !strings.HasSuffix(shortcutTargetFilePath, ".lnk") {
		return "", ErrWinOsExeShortcutCreationInput
	}

	var linkCreationCommandString = fmt.Sprintf(`Set oWS = WScript.CreateObject("WScript.Shell")
		Set oLink = oWS.CreateShortcut("%v")
			oLink.TargetPath = "%v"
			oLink.Arguments = "%v"
			oLink.WindowStyle = "1"   
		'	oLink.WorkingDirectory = ""
		'	oLink.Description = ""
		'	oLink.IconLocation = ""
		'	oLink.HotKey = ""
		oLink.Save`,

		shortcutTargetFilePath,

		shortcutSourceFilePath,

		// NOTE: in VBS double quotes escaping is possible by doubling the quotes themselves
		// https://stackoverflow.com/questions/42307239/vbs-to-create-desktop-shortcut-with-runas-arguments
		strings.Replace(shortcutArguments, "\"", "\"\"", -1),
	)

	return linkCreationCommandString, nil
}
