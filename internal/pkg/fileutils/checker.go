package fileutils

import (
	"os"
)

func IsDirExists(dir string) bool {
	_, err := os.Stat(dir)

	return err != nil
}
