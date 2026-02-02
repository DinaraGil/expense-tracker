package storage

import (
	"os"
	"path/filepath"
)

const dir = "storage"
const filename = "tasks.json"

func GetStoragePath() (string, error) {
	exePath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(exePath, "internal", dir, filename), nil
}
