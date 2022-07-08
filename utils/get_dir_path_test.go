package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDirPath(t *testing.T) {
	_, err := GetDirPath("/doesnt/exist")

	assert.NotNil(
		t,

		err,

		"Should return an error when the given path doesn't exist",
	)

	dir, _ := GetDirPath("../data")

	assert.True(
		t,

		strings.HasSuffix(dir, "/data"),

		"Should be able to return a given directory path as is",
	)

	file, _ := GetDirPath("../data/executable")

	assert.True(
		t,

		strings.HasSuffix(file, "/data/"),

		"Should be able to return the directory path of a given file path",
	)
}
