# retroarch-links-generator

Generate links to RetroArch playlists entries and launch them like normal games.

## Usage

[Download](https://github.com/memob0x/retroarch-links-generator/releases) the binary file corresponding to your OS environment and launch it from command line.

The **first argument** is the path to the RetroArch executable file.

The **second argument** is the game links destination folder path.

The **third argument**, _optional_, is the RetroArch playlist folder path; on windows this is not necessary since that folder is in the same folder of the executable file.

### Windows
```console
retroarch-links-generator-windows-amd64.exe D:\SteamLibrary\steamapps\common\RetroArch\retroarch.exe C:\retroarch-launchers
``` 

### Linux
```console
retroarch-links-generator-linux-amd64 /usr/bin/retroarch ~/retroarch-launchers ~/Documents/RetroArch
``` 

### Mac
```console
retroarch-links-generator-darwin-amd64 /Applications/RetroArch ~/retroarch-launchers ~/Documents/RetroArch
``` 