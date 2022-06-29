package utils

func AppendUniquePlaylist(playlistsInput []PlaylistInfo, playlistsToBeAdded PlaylistInfo) []PlaylistInfo {
	var playlistsOutput []PlaylistInfo = playlistsInput

	for _, v := range playlistsInput {
		if playlistsToBeAdded.Path == v.Path {
			return playlistsOutput
		}
	}

	return append(playlistsInput, playlistsToBeAdded)
}
