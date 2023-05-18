package tmplcreator

import (
	"html/template"
	"os"
	"path"
)

type CommonCreator struct {
	projectDirectory string
	appName          string
	moduleName       string
}

func NewCommonCreator(projectDirectory string, appName, moduleName string) *CommonCreator {
	return &CommonCreator{
		projectDirectory: projectDirectory,
		appName:          appName,
		moduleName:       moduleName,
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
	createMainFolderErr := os.MkdirAll(path.Join(c.projectDirectory, "cmd", c.appName), defaultFolderPerm)
	if createMainFolderErr != nil && !os.IsExist(createMainFolderErr) {
		return createMainFolderErr
	}

	createInternalFolderErr := os.MkdirAll(path.Join(c.projectDirectory, "internal"), defaultFolderPerm)
	if createInternalFolderErr != nil && !os.IsExist(createInternalFolderErr) {
		return createInternalFolderErr
	}

	createPkgFolderErr := os.MkdirAll(path.Join(c.projectDirectory, "pkg"), defaultFolderPerm)
	if createPkgFolderErr != nil && !os.IsExist(createMainFolderErr) {
		return createPkgFolderErr
	}

	createLogsFolderErr := os.MkdirAll(path.Join(c.projectDirectory, "logs"), defaultFolderPerm)
	if createLogsFolderErr != nil && !os.IsExist(createLogsFolderErr) {
		return createLogsFolderErr
	}

	createMigrationsFolderErr := os.MkdirAll(path.Join(c.projectDirectory, "migrations"), defaultFolderPerm)
	if createMigrationsFolderErr != nil && !os.IsExist(createMigrationsFolderErr) {
		return createMigrationsFolderErr
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

	if logErr := c.createLogFile(); logErr != nil {
		return logErr
	}

	if mainErr := c.createMainFile(); mainErr != nil {
		return mainErr
	}

	return nil
}

func (c *CommonCreator) createGitignoreFile() error {
	f, createErr := os.Create(path.Join(c.projectDirectory, ".gitignore"))
	if createErr != nil {
		return createErr
	}
	defer f.Close()

	tmpl, parseTmplErr := template.ParseFS(templatesFolder, gitignoreTemplatePath)
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
	var data = map[string]string{"AppName": c.appName}

	tmpl, parseTmplErr := template.ParseFS(templatesFolder, makefileTemplatePath)
	if parseTmplErr != nil {
		return parseTmplErr
	}

	if executeErr := tmpl.Execute(f, data); executeErr != nil {
		return executeErr
	}

	return nil
}

func (c *CommonCreator) createLogFile() error {
	f, createErr := os.Create(path.Join(c.projectDirectory, "logs", "log.log"))
	if createErr != nil {
		return createErr
	}
	defer f.Close()

	return nil
}

func (c *CommonCreator) createMainFile() error {
	f, createErr := os.Create(path.Join(c.projectDirectory, "cmd", c.appName, "main.go"))
	if createErr != nil {
		return createErr
	}
	defer f.Close()

	var data = map[string]string{"ModuleName": c.moduleName}

	tmpl, parseTmplErr := template.ParseFS(templatesFolder, mainTemplatePath)
	if parseTmplErr != nil {
		return parseTmplErr
	}

	if executeErr := tmpl.Execute(f, data); executeErr != nil {
		return executeErr
	}

	return nil
}
