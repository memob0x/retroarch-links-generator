package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVbsScriptShortcutCreatorContent(t *testing.T) {
	content, err := GetVbsScriptShortcutCreatorContent(
		"foo-source",

		"./foo-link.lnk",

		"foo-arguments",
	)

	assert.Equal(
		t,

		content,

		"",

		"Should return an error if a source without .exe extension is passed",
	)

	assert.Equal(
		t,

		err,

		ErrWinOsExeShortcutCreationInput,

		"Should return an error if a source without .exe extension is passed",
	)

	content, err = GetVbsScriptShortcutCreatorContent(
		"foo-source",

		"./foo-link.lnk",

		"foo-arguments",
	)

	assert.Equal(
		t,

		content,

		"",

		"Should return an error if a source without .exe extension is passed",
	)

	assert.Equal(
		t,

		err,

		ErrWinOsExeShortcutCreationInput,

		"Should return an error if a target without .lnk extension is passed",
	)

	content, err = GetVbsScriptShortcutCreatorContent(
		"foo-source.exe",

		"./foo-link.lnk",

		"foo-arguments",
	)

	assert.Equal(
		t,

		content,

		"Set oWS = WScript.CreateObject(\"WScript.Shell\")\n\t\tSet oLink = oWS.CreateShortcut(\"./foo-link.lnk\")\n\t\t\toLink.TargetPath = \"foo-source.exe\"\n\t\t\toLink.Arguments = \"foo-arguments\"\n\t\t\toLink.WindowStyle = \"1\"   \n\t\t'\toLink.WorkingDirectory = \"\"\n\t\t'\toLink.Description = \"\"\n\t\t'\toLink.IconLocation = \"\"\n\t\t'\toLink.HotKey = \"\"\n\t\toLink.Save",

		"Should return an error if a source without .exe extension is passed",
	)

	assert.NotEqual(
		t,

		err,

		ErrWinOsExeShortcutCreationInput,

		"Should not return an error if valid *.exe source and *.lnk target are passed",
	)
}
