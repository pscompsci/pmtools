package pmtools

import (
	"errors"
	"os"
	"path"
)

// DirectoryPath generates the path string for a directory
func DirectoryPath(p *Project, folder string) (string, error) {
	if p.BaseFolder == "" {
		return "", errors.New("base folder is not set")
	}
	fi, err := os.Stat(p.BaseFolder)
	if err != nil {
		return "", errors.New("base folder is not accessible")
	}
	if !fi.Mode().IsDir() {
		return "", errors.New("base folder is not a directory")
	}
	return path.Join(p.BaseFolder, folder), nil
}

// CreateDirectory creates a directory with the desired path,
// if it doesn't exist
func CreateDirectory(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.Mkdir(path, os.ModeDir)
		return nil
	}
	return err
}
