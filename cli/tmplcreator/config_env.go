package tmplcreator

import (
	"os"
	"path"
)

type EnvConfigCreator struct {
	projectDirectory string
}

func NewEnvConfigCreator(projectDirectory string) *EnvConfigCreator {
	return &EnvConfigCreator{projectDirectory: projectDirectory}
}

func (e *EnvConfigCreator) Create() error {
	createConfigFolderErr := os.MkdirAll(path.Join(e.projectDirectory, "configs"), defaultFolderPerm)
	if createConfigFolderErr != nil && !os.IsExist(createConfigFolderErr) {
		return createConfigFolderErr
	}

	if createEnvFileErr := e.writeToEnvFile(); createEnvFileErr != nil {
		return createEnvFileErr
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
