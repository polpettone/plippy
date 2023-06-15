package cmd

import (
	"os"
	"path/filepath"
)

const plippyFilePath = ".config/plippy/contents.yaml"

func plippyFile() (string, error) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(homeDir, plippyFilePath)
	return filePath, nil
}
