package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitWithRegex(t *testing.T) {
	assert.Equal(
		t,

		[]string{"foobar", "1", "2", "3"},

		SplitStringWithRegex("foobar\\1/2/3", `\\|/`),

		" ",
	)
}
