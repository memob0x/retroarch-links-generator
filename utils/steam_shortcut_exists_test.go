package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSteamShortcutExists(t *testing.T) {
	var shortcuts = []SteamShortcut{
		{
			AppName: "foobar",
		},
		{
			AppId: 123,
		},
	}

	assert.True(
		t,

		SteamShortcutExists(
			shortcuts,

			SteamShortcut{
				AppId: 123,
			},
		),

		"Should be able to return true when a given shortcut has same AppId field value of another shortcut in the given shortcuts list",
	)

	assert.True(
		t,

		SteamShortcutExists(
			shortcuts,

			SteamShortcut{
				AppName: "foobar",
			},
		),

		"Should be able to return true when a given shortcut has same AppName field value of another shortcut in the given shortcuts list",
	)

	assert.False(
		t,

		SteamShortcutExists(
			shortcuts,

			SteamShortcut{
				AppId: 345,

				AppName: "fizzbuz",
			},
		),

		"Should be able to return false when a given shortcut has not AppId nor AppName field value of another shortcut in the given shortcuts list",
	)
}
