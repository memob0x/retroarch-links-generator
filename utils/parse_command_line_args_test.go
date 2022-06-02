package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommandLineArgs(t *testing.T) {
	a, b := ParseCommandLineArgs([]string{"foo", "bar", "biz", "abc", "extraaaa"})

	assert.Equal(
		t,

		"bar",

		a,

		"should be able to extract second value from the argument list",
	)

	assert.Equal(
		t,

		"biz",

		b,

		"should be able to extract third value from the argument list",
	)
}
