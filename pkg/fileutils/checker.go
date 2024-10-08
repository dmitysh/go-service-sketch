package fileutils

import (
	"os"
)

func IsDirExists(dir string) bool {
	_, err := os.Stat(dir)
	if err != nil {
		return false
	}

	return true
}
