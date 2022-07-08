package utils

import (
	"hash/crc32"
)

func GenerateSteamAppID(appExecutablePath string) uint32 {
	return crc32.ChecksumIEEE([]byte(appExecutablePath))
}
