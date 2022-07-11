package utils

import (
	"io/ioutil"
	"os"
)

// Credits: https://stackoverflow.com/questions/32438204/create-a-windows-shortcut-lnk-in-go

func CreateWinOsExeShortcut(
	shortcutSourceFilePath string,

	shortcutTargetFilePath string,

	shortcutArguments string,
) error {
	var linkCreationCommandString, contentError = GetVbsScriptShortcutCreatorContent(
		shortcutSourceFilePath,

		shortcutTargetFilePath,

		shortcutArguments,
	)

	if contentError != nil {
		return contentError
	}

	var vbsScriptPath string = "./shortcut-creator.vbs"

	var fileError = ioutil.WriteFile(vbsScriptPath, []byte(linkCreationCommandString), 0777)

	if fileError != nil {
		return fileError
	}

	var executionError = ExecuteVbsScript(vbsScriptPath)

	fileError = os.Remove(vbsScriptPath)

	if executionError != nil {
		return executionError
	}

	return fileError
}
