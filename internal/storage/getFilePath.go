package storage

import (
	"os"
	"path/filepath"
)

const dir = "storage"
const filename = "tasks.json"

type Filename string

const (
	TempFile  Filename = "temp.json"
	ConstFile Filename = "tasks.json"
)

func GetStoragePath(filename Filename) (string, error) {
	exePath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(exePath, "internal", dir, string(filename)), nil
}
