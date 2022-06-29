package utils

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
)

type FileInfo struct {
	Path string

	FsFile fs.FileInfo
}

func GetAllFilesRecursively(fileStatsInput []FileInfo) ([]FileInfo, error) {
	var filesStatsOutput []FileInfo

	for _, fileStat := range fileStatsInput {
		fileStatInfo := fileStat.FsFile

		if !fileStatInfo.IsDir() {
			filesStatsOutput = append(filesStatsOutput, fileStat)

			continue
		}

		dirPath, err := filepath.Abs(fileStat.Path)

		if err != nil {
			return nil, err
		}

		itemsInDir, err := ioutil.ReadDir(dirPath)

		if err != nil {
			return nil, err
		}

		var infosWithDirPaths []FileInfo

		for _, fileInfo := range itemsInDir {
			filePath, err := filepath.Abs(dirPath + "/" + fileInfo.Name())

			if err != nil {
				return nil, err
			}

			infosWithDirPaths = append(infosWithDirPaths, FileInfo{
				Path: filePath,

				FsFile: fileInfo,
			})
		}

		filesInDir, err := GetAllFilesRecursively(infosWithDirPaths)

		if err != nil {
			return nil, err
		}

		filesStatsOutput = append(filesStatsOutput, filesInDir...)
	}

	return filesStatsOutput, nil
}
