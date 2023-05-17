package config

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
	err := os.MkdirAll(path.Join(e.projectDirectory, "configs"), 0666)
	if err != nil && !os.IsExist(err) {
		return err
	}

	return nil
}
