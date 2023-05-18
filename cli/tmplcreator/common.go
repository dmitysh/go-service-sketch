package tmplcreator

import (
	"html/template"
	"os"
	"path"
)

const (
	makefileTemplatePath  = "cli/tmplcreator/templates/makefile.txt"
	gitignoreTemplatePath = "cli/tmplcreator/templates/gitignore.txt"
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

	tmpl, parseTmplErr := template.ParseFiles(gitignoreTemplatePath)
	if parseTmplErr != nil {
		return parseTmplErr
	}

	if executeErr := tmpl.Execute(f, nil); executeErr != nil {
		return executeErr
	}

	return nil
}

func (c *CommonCreator) createMakefileFile() error {
	f, createErr := os.Create(path.Join(c.projectDirectory, "Makefile"))
	if createErr != nil {
		return createErr
	}
	defer f.Close()
	data := makefileTemplateData{AppName: c.appName}

	tmpl, parseTmplErr := template.ParseFiles(makefileTemplatePath)
	if parseTmplErr != nil {
		return parseTmplErr
	}

	if executeErr := tmpl.Execute(f, data); executeErr != nil {
		return executeErr
	}

	return nil
}
