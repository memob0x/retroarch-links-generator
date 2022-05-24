package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetValidWinOsFilenameTest(t *testing.T) {
	assert.Equal(t, "-aaa-aaa-aaa-aaa-aaa-aaa-aaa-aaa-", GetValidWinOsFilename("\aaa/aaa:aaa*aaa?aaa\"aaa<aaa>aaa|"), "should be able to format float leaving 1 decimal")
}
