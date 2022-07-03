package utils

import "hash/crc32"

// thanks https://github.com/Wakeful-Cloud/aluminum/blob/a246c2538c8f49be5fd9a91e9535870fa9f3d1d1/steam/steam.go
func GenerateSteamAppID(appExecutablePath string) uint32 {
	checksum := crc32.ChecksumIEEE([]byte(appExecutablePath))

	a := checksum | 0x80000000
	b := uint64(a) << 32
	c := b | 0x02000000

	return uint32(c)
}
