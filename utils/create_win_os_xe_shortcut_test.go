package utils

import (
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWinOsShortcut(t *testing.T) {
	err := CreateWinOsExeShortcut(
		"foo-source",

		"./foo-link.lnk",

		"foo-arguments",
	)

	assert.Equal(
		t,

		err,

		WinOsExeShortcutCreationInputError,

		"Should return an error if a source without .exe extension is passed",
	)

	err = CreateWinOsExeShortcut(
		"foo-source",

		"./foo-link.lnk",

		"foo-arguments",
	)

	assert.Equal(
		t,

		err,

		WinOsExeShortcutCreationInputError,

		"Should return an error if a target without .lnk extension is passed",
	)

	err = CreateWinOsExeShortcut(
		"foo-source.exe",

		"./foo-link.lnk",

		"foo-arguments",
	)

	assert.NotEqual(
		t,

		err,

		WinOsExeShortcutCreationInputError,

		"Should not return an error if valid *.exe source and *.lnk target are passed",
	)

	if runtime.GOOS != "windows" {
		// TODO: add non-windows tests

		return
	}

	// TODO: check if it effectively passes on windows

	s, _ := os.Stat("./foo-link.lnk")

	assert.Greater(
		t,

		s.Size(),

		int64(0),

		"Should be able to write a Windows os Lnk file",
	)

	os.Remove("./foo-link.lnk")
}
