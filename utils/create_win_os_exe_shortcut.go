package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Credits: https://stackoverflow.com/questions/32438204/create-a-windows-shortcut-lnk-in-go

var WinOsExeShortcutCreationInputError = errors.New("Invalid executable paths provided")

func CreateWinOsExeShortcut(
	shortcutSourceFilePath string,

	shortcutTargetFilePath string,

	shortcutArguments string,
) error {
	if !strings.HasSuffix(shortcutSourceFilePath, ".exe") || !strings.HasSuffix(shortcutTargetFilePath, ".lnk") {
		return WinOsExeShortcutCreationInputError
	}

	var scriptTxt bytes.Buffer

	scriptTxt.WriteString(fmt.Sprintf(`Set oWS = WScript.CreateObject("WScript.Shell")
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
	))

	var filename string = "shortcut-creator.vbs"

	ioutil.WriteFile(filename, scriptTxt.Bytes(), 0777)

	var cmd *exec.Cmd = exec.Command("wscript", filename)

	var err error = cmd.Run()

	cmd.Wait()

	os.Remove(filename)

	return err
}
