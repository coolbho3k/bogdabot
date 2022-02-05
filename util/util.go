package util

import (
	"os"
)

// IsDirectory checks if a file path is a directory.
func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}
