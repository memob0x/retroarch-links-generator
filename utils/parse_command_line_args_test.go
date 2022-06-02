package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommandLineArgs(t *testing.T) {
	a, b, c, _ := ParseCommandLineArgs([]string{"command", "/path/to/executable", "/path/to/links/output", "/path/to/playlists", "extra", "args"})

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

	d, e, f, _ := ParseCommandLineArgs([]string{"command", "/path/to/executable", "/path/to/links/output"})

	assert.Contains(
		t,

		d,

		"/path/to/executable",

		"should be able to extract second value from the argument list",
	)

	assert.Contains(
		t,

		"/path/to/links/output",

		e,

		"should be able to extract third value from the argument list",
	)

	assert.Contains(
		t,

		"/path/to/playlists",

		f,

		"should be able to extract fourth value from the argument list",
	)
}
