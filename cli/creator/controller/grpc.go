package controller

import (
	"os"
	"path"
)

type GRPCControllerCreator struct {
	projectDirectory string
}

func NewGRPCControllerCreator(projectDirectory string) *GRPCControllerCreator {
	return &GRPCControllerCreator{projectDirectory: projectDirectory}
}

func (g *GRPCControllerCreator) Create() error {
	err := os.MkdirAll(path.Join(g.projectDirectory, "server"), 0666)
	if err != nil && !os.IsExist(err) {
		return err
	}

	return nil
}
