package utils

func AppendUniquePlaylists(playlistsInput []PlaylistInfo, playlistsToBeAdded []PlaylistInfo) []PlaylistInfo {
	var playlistsOutput []PlaylistInfo = playlistsInput

	for _, playlistTobeAdded := range playlistsToBeAdded {
		playlistsOutput = AppendUniquePlaylist(playlistsOutput, playlistTobeAdded)
	}

	return playlistsOutput
}
