package tmplcreator

import (
	"os"
	"path"
	"text/template"
)

type EnvConfigCreator struct {
	projectDirectory string
	appName          string
}

func NewEnvConfigCreator(projectDirectory string, appName string) *EnvConfigCreator {
	return &EnvConfigCreator{projectDirectory: projectDirectory, appName: appName}
}

func (e *EnvConfigCreator) Create() error {
	createConfigsFolderErr := os.MkdirAll(path.Join(e.projectDirectory, "configs"), defaultFolderPerm)
	if createConfigsFolderErr != nil && !os.IsExist(createConfigsFolderErr) {
		return createConfigsFolderErr
	}

	if createEnvFileErr := e.writeToEnvFile(); createEnvFileErr != nil {
		return createEnvFileErr
	}

	if createUtilsFileErr := e.createUtilsFile(); createUtilsFileErr != nil {
		return createUtilsFileErr
	}

	if addConfigErr := e.addConfigInfoToMainFile(); addConfigErr != nil {
		return addConfigErr
	}

	return nil
}

func (e *EnvConfigCreator) createUtilsFile() error {
	createConfigFolderErr := os.MkdirAll(path.Join(e.projectDirectory, "configs", "config"), defaultFolderPerm)
	if createConfigFolderErr != nil && !os.IsExist(createConfigFolderErr) {
		return createConfigFolderErr
	}

	f, createErr := os.Create(path.Join(e.projectDirectory, "configs", "config", "utils.go"))
	if createErr != nil {
		return createErr
	}
	defer f.Close()

	tmpl, parseTmplErr := template.ParseFS(templatesFolder, envUtilsTemplatePath)
	if parseTmplErr != nil {
		return parseTmplErr
	}

	if executeErr := tmpl.Execute(f, nil); executeErr != nil {
		return executeErr
	}

	return nil
}

func (e *EnvConfigCreator) writeToEnvFile() error {
	f, openErr := os.OpenFile(path.Join(e.projectDirectory, "configs", "app.env"),
		os.O_CREATE|os.O_APPEND|os.O_WRONLY, defaultFilePerm)
	if openErr != nil {
		return openErr
	}
	defer f.Close()

	_, writeErr := f.WriteString(`SERVER_HOST=0.0.0.0
SERVER_PORT=8940

`)
	if writeErr != nil {
		return writeErr
	}

	return nil
}

func (e *EnvConfigCreator) addConfigInfoToMainFile() error {
	var data = map[string]string{"ConfigExt": "env", "ConfigType": "Env"}

	tmpl, parseTmplErr := template.ParseFiles(path.Join(e.projectDirectory, "cmd", e.appName, "main.go"))
	if parseTmplErr != nil {
		return parseTmplErr
	}

	f, createErr := os.Create(path.Join(e.projectDirectory, "cmd", e.appName, "main.go"))
	if createErr != nil {
		return createErr
	}
	defer f.Close()

	if executeErr := tmpl.Execute(f, data); executeErr != nil {
		return executeErr
	}

	return nil
}
