package utils

import (
	"os"
	"path/filepath"
)

func ConvertToFullPath(relativePath string) (string, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	executableDir := filepath.Dir(executablePath)
	configFilePath := filepath.Join(executableDir, relativePath)

	return configFilePath, nil
}
