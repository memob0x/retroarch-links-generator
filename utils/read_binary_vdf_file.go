package utils

import (
	"errors"
	"os"

	"github.com/wakeful-cloud/vdf"
)

var ErrorVdfNotExists = errors.New("No vdf file.")

func ReadBinaryVdfFile(vdfFilePath string) (vdf.Map, error) {
	var vdfMap vdf.Map

	_, err := os.Stat(vdfFilePath)

	if os.IsNotExist(err) {
		return vdfMap, ErrorVdfNotExists
	}

	bytes, err := os.ReadFile(vdfFilePath)

	if err != nil {
		return vdfMap, err
	}

	vdfMap, err = vdf.ReadVdf(bytes)

	if err != nil {
		return vdfMap, err
	}

	return vdfMap, nil
}
