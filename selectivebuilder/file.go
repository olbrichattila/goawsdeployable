package main

import (
	"os"
	"path/filepath"
)

type fileInfo struct {
	path     string
	relPath  string
	fileInfo os.FileInfo
}

func readDir(sourceFolder string) (*[]fileInfo, error) {
	var files []fileInfo

	err := filepath.Walk(sourceFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			relPath, err := filepath.Rel(sourceFolder, path)
			if err != nil {
				return err
			}
			dir, _ := filepath.Split(relPath)

			files = append(files, fileInfo{fileInfo: info, path: path, relPath: dir})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &files, nil
}

func mkDir(destinationFolder string) error {
	err := os.MkdirAll(destinationFolder, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
