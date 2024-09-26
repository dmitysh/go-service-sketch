package temple

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
	if createControllerFoldersErr := g.createControllerFolders(); createControllerFoldersErr != nil {
		return createControllerFoldersErr
	}

	if createControllerFilesErr := g.createControllerFiles(); createControllerFilesErr != nil {
		return createControllerFilesErr
	}

	return nil
}

func (g *GRPCControllerCreator) createControllerFolders() error {
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

func (g *GRPCControllerCreator) createControllerFiles() error {
	f, createApiKeepFileErr := os.Create(path.Join(g.projectDirectory, "api", g.appName, "keep.proto"))
	if createApiKeepFileErr != nil {
		return createApiKeepFileErr
	}
	defer f.Close()

	return nil
}
