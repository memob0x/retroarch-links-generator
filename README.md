# retroarch-links-generator

Generate links to RetroArch playlists entries and launch them like normal games.

![how-it-works](./docs/how-it-works.gif)

## Basic usage

[Download](https://github.com/memob0x/retroarch-links-generator/releases) the binary file corresponding to your OS environment and launch it from command line with the following arguments:

1. The path to the RetroArch executable file.
2. The game links destination folder path.
3. The RetroArch playlist folder path.

```console
retroarch-links-generator /the/retroarch/executable /the/destination/folder /the/retroarch/playlists
```

## Advanced

### Only parse certain playlists
```console
retroarch-links-generator /the/retroarch/executable /the/destination/folder /a/retroarch/playlist-1.lpl,/a/retroarch/playlist-2.lpl
``` 

### Steam shortcuts (soon)
```console
retroarch-links-generator /the/retroarch/executable /the/steam/libraryfolder.vdf
``` 