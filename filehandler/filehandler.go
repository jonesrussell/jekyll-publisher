package filehandler

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

type FileHandlerInterface interface {
	GetFilenames(sitePath string, dirName string) ([]string, error)
}

type FileHandler struct {
}

// Ensure Menu implements MenuInterface
var _ FileHandlerInterface = &FileHandler{}

func (fh FileHandler) GetFilenames(sitePath string, dirName string) ([]string, error) {
	var filenames []string
	dirPath := filepath.Join(sitePath, dirName)
	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			filenames = append(filenames, filepath.Base(path))
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error reading %s: %v", dirName, err)
	}
	return filenames, nil
}
