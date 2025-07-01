package fileutils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyFile(srcPath, dstPath string) error {
	src, err := os.Open(srcPath) //nolint:gosec
	if err != nil {
		return fmt.Errorf("can't open file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(dstPath) //nolint:gosec
	if err != nil {
		return fmt.Errorf("can't create file: %w", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("can't make copy: %w", err)
	}

	return nil
}

func CopyDir(srcPath, dstPath string) error {
	entries, err := os.ReadDir(srcPath)
	if err != nil {
		return fmt.Errorf("can't read dir: %w", err)
	}

	err = os.MkdirAll(dstPath, 0750)
	if err != nil {
		return fmt.Errorf("can't make dir: %w", err)
	}

	for _, entry := range entries {
		curSrcPath := filepath.Join(srcPath, entry.Name())
		curDstPath := filepath.Join(dstPath, entry.Name())

		if entry.IsDir() {
			if err := CopyDir(curSrcPath, curDstPath); err != nil {
				return err
			}
		} else {
			if err := CopyFile(curSrcPath, curDstPath); err != nil {
				return err
			}
		}
	}

	return nil
}
