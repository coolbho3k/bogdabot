package commands

import (
	"bogdabot/util"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Store interface {
	// GetResponseByPath gets the response of a slash command by its path.
	GetResponseByPath(command string) (string, error)
}

type store struct {
	responseMap map[string]string
}

func New() (Store, error) {
	responseMap := make(map[string]string)

	var files []string
	root := "data/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if isDir, err := util.IsDirectory(file); isDir {
			continue
		} else if err != nil {
			return nil, err
		}

		if filepath.Ext(file) != ".json" {
			continue
		}

		data, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		responseMap[filepath.Base(file[:len(file)-len(filepath.Ext(file))])] = string(data)
	}

	return &store{
		responseMap,
	}, nil
}

func (s *store) GetResponseByPath(command string) (string, error) {
	if response, ok := s.responseMap[command]; ok {
		return response, nil
	}
	return "", errors.New("command not found")
}

