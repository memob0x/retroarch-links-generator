package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommandLineArgs(t *testing.T) {
	a, b, c, cbool, _ := ParseCommandLineArgs([]string{"command", "/path/to/executable", "/path/to/links/output", "/path/to/playlists", "extra", "args"})

	assert.Contains(
		t,

		a,

		"/path/to/executable",

		"should be able to extract second value from the argument list",
	)

	assert.Contains(
		t,

		b,

		"/path/to/links/output",

		"should be able to extract third value from the argument list",
	)

	assert.Contains(
		t,

		c,

		"/path/to/playlists",

		"should be able to extract fourth value from the argument list",
	)

	assert.False(
		t,

		cbool,

		"should be able to detect a not given vdf output file",
	)

	d, e, f, fbool, _ := ParseCommandLineArgs([]string{"command", "/path/to/executable", "/path/to/links/output.vdf"})

	assert.Contains(
		t,

		d,

		"/path/to/executable",

		"should be able to extract second value from the argument list",
	)

	assert.Contains(
		t,

		"/path/to/links/output.vdf",

		e,

		"should be able to extract third value from the argument list",
	)

	assert.Contains(
		t,

		"/path/to/playlists",

		f,

		"should be able to extract fourth value from the argument list",
	)

	assert.True(
		t,

		fbool,

		"should be able to detect a given vdf output file",
	)
}
