package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSteamAppID(t *testing.T) {
	assert.NotEqual(
		t,

		GenerateSteamAppID("abc.exe --foobar"),

		GenerateSteamAppID("abc.exe --fizzbuzz"),

		"a",
	)
}
