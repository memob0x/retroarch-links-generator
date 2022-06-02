package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLinkFileInfosFromPlaylistItem(t *testing.T) {
	result, _ := GetLinkFileInfosFromPlaylistItem(RetroArchPlaylistItem{
		RomPath: "/path/to/rom",

		CorePath: "/path/to/core",

		Label: "Foo Game Bar",
	}, "/here", "/there")

	assert.Equal(
		t,

		"/here\\retroarch.exe --load-menu-on-error -L /path/to/core /path/to/rom",

		result.Content,

		"aa",
	)

	assert.Equal(
		t,

		"/there\\Foo Game Bar.bat",

		result.Path,

		"aa",
	)
}
