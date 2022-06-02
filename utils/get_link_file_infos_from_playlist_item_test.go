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
	}, "/here/executable.xyz", "/there")

	assert.Contains(
		t,

		result.Content,

		"/here/executable.xyz --load-menu-on-error -L /path/to/core /path/to/rom",

		"Should be able to return an object with valid command content string",
	)

	assert.Contains(
		t,

		result.Path,

		"/there/Foo Game Bar",

		"Should be able to return an object with the future link path",
	)
}
