package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValidWinOsFilename(t *testing.T) {
	assert.Equal(
		t,

		"-aaa-aaa-aaa-aaa-aaa-aaa-aaa-aaa-",

		GetValidWinOsFilename("\\aaa/aaa:aaa*aaa?aaa\"aaa<aaa>aaa|"),

		"should be able to format strings making compatible for a Windows OS environment",
	)
}
