package tmplcreator

import (
	"fmt"
	"os"
	"path"
)

type CommonCreator struct {
	projectDirectory string
	appName          string
}

func NewCommonCreator(projectDirectory string, appName string) *CommonCreator {
	return &CommonCreator{
		projectDirectory: projectDirectory,
		appName:          appName,
	}
}

func (c *CommonCreator) Create() error {
	if createCommonFoldersErr := c.createCommonFolders(); createCommonFoldersErr != nil {
		return createCommonFoldersErr
	}
	if createCommonFilesErr := c.createCommonFiles(); createCommonFilesErr != nil {
		return createCommonFilesErr
	}

	return nil
}

func (c *CommonCreator) createCommonFolders() error {
	createMainFolderErr := os.MkdirAll(path.Join(c.projectDirectory, "cmd", c.appName), 0666)
	if createMainFolderErr != nil && !os.IsExist(createMainFolderErr) {
		return createMainFolderErr
	}
	createInternalFolderErr := os.MkdirAll(path.Join(c.projectDirectory, "internal"), 0666)
	if createInternalFolderErr != nil && !os.IsExist(createInternalFolderErr) {
		return createInternalFolderErr
	}
	createPkgFolderErr := os.MkdirAll(path.Join(c.projectDirectory, "pkg"), 0666)
	if createPkgFolderErr != nil && !os.IsExist(createMainFolderErr) {
		return createPkgFolderErr
	}

	return nil
}

func (c *CommonCreator) createCommonFiles() error {
	if gitignoreErr := c.createGitignoreFile(); gitignoreErr != nil {
		return gitignoreErr
	}
	if makefileErr := c.createMakefileFile(); makefileErr != nil {
		return makefileErr
	}

	return nil
}

func (c *CommonCreator) createGitignoreFile() error {
	f, createErr := os.Create(path.Join(c.projectDirectory, ".gitignore"))
	if createErr != nil {
		return createErr
	}
	defer f.Close()

	_, writerErr := f.WriteString(".idea\n")
	if writerErr != nil {
		return writerErr
	}

	return nil
}

func (c *CommonCreator) createMakefileFile() error {
	f, createErr := os.Create(path.Join(c.projectDirectory, "Makefile"))
	if createErr != nil {
		return createErr
	}
	defer f.Close()

	_, writerErr := f.WriteString(fmt.Sprintf(`.PHONY: run
run:
	go run cmd/%s/main.go
`, c.appName))
	if writerErr != nil {
		return writerErr
	}

	return nil
}
