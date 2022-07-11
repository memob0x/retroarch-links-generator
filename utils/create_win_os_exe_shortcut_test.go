package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testCreateWinOsShortcutWindows(t *testing.T) {
	if !IS_OS_WINDOWS {
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

func testCreateWinOsShortcutNonWindows(t *testing.T) {
	if IS_OS_WINDOWS {
		return
	}

	// TODO: add non-windows tests
}

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

	testCreateWinOsShortcutNonWindows(t)

	testCreateWinOsShortcutWindows(t)
}
