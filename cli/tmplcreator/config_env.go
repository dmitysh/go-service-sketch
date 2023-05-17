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
	createConfigFolderErr := os.MkdirAll(path.Join(e.projectDirectory, "configs"), 0666)
	if createConfigFolderErr != nil && !os.IsExist(createConfigFolderErr) {
		return createConfigFolderErr
	}

	return nil
}
