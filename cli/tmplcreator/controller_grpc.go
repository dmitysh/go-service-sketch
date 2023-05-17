package tmplcreator

import (
	"os"
	"path"
)

type GRPCControllerCreator struct {
	projectDirectory string
	appName          string
}

func NewGRPCControllerCreator(projectDirectory string, appName string) *GRPCControllerCreator {
	return &GRPCControllerCreator{
		projectDirectory: projectDirectory,
		appName:          appName,
	}
}

func (g *GRPCControllerCreator) Create() error {
	createServerFolderErr := os.MkdirAll(path.Join(g.projectDirectory, "server"), 0666)
	if createServerFolderErr != nil && !os.IsExist(createServerFolderErr) {
		return createServerFolderErr
	}

	createAPIFolderErr := os.MkdirAll(path.Join(g.projectDirectory, "api", g.appName), 0666)
	if createAPIFolderErr != nil && !os.IsExist(createAPIFolderErr) {
		return createAPIFolderErr
	}

	return nil
}
