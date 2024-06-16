package main

import (
	"io"
	"os"
	"path/filepath"
)

type fileInfo struct {
	path     string
	relPath  string
	fileInfo os.FileInfo
}

func copyDir(src string, dest string) error {
	files, err := readDir(src)
	if err != nil {
		return err
	}
	mkDir(dest)

	for _, f := range *files {
		err := copyFile(f.path, dest+f.relPath+"/"+f.fileInfo.Name())
		if err != nil {
			return err
		}
	}

	return nil

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

func copyFile(sourceFileName, destFileName string) error {

	sourceFile, err := os.Open(sourceFileName)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(destFileName)
	if err != nil {
		return nil
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	err = destFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

func rmDir(dirPath string) error {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return os.RemoveAll(path)
	})

	if err != nil {
		return err
	}

	if err := os.RemoveAll(dirPath); err != nil {
		return err
	}

	return nil
}
