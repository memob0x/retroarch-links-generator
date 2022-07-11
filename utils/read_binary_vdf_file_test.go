package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wakeful-cloud/vdf"
)

func TestReadBinaryVdfFile(t *testing.T) {
	vdfMap := vdf.Map{
		"test": vdf.Map{
			"0": "1",

			"1": uint32(2),
		},
	}

	vdfMapBytes, _ := vdf.WriteVdf(vdfMap)

	os.WriteFile("./binary-file.vdf", vdfMapBytes, 0655)

	vdfMapRead, _ := ReadBinaryVdfFile("./binary-file.vdf")

	assert.Equal(
		t,

		vdfMap,

		vdfMapRead,

		"Should be able to read a binary vdf file",
	)

	os.Remove("./binary-file.vdf")
}
