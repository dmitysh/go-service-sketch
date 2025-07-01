package temple

import (
	"os"
	"path"
)

type CommonCreatorParams struct {
	ProjectDirectory string
	AppName          string
	ModuleName       string
	GoVersion        string
}

type CommonCreator struct {
	params         CommonCreatorParams
	fileCreators   []func() error
	folderCreators []func() error
}

func NewCommonCreator(params CommonCreatorParams) *CommonCreator {
	c := &CommonCreator{
		params: params,
	}

	c.fileCreators = []func() error{
		c.createGitignoreFile,
		c.createMakefileFile,
		c.createDockerfileFile,
		c.createDockerComposeFile,
		c.createLocalEnvFile,
		c.createProdEnvFile,
		c.createDockerEnvFile,
		c.createLocalConfigFile,
		c.createProdConfigFile,
		c.createDockerConfigFile,
		c.createGoModFile,
		c.createMainFile,
		c.createAppFile,
		c.createConfigFile,
	}

	c.folderCreators = []func() error{
		c.createServiceFolder,
		c.createDomainFolder,
		c.createUsecaseFolder,
	}

	return c
}

func (c *CommonCreator) Create() error {
	if err := createFolders(c.folderCreators); err != nil {
		return err
	}

	if err := createFiles(c.fileCreators); err != nil {
		return err
	}

	return nil
}

func (c *CommonCreator) createServiceFolder() error {
	err := os.MkdirAll(path.Join(c.params.ProjectDirectory, "internal", "service"), defaultFolderPerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	return nil
}

func (c *CommonCreator) createDomainFolder() error {
	err := os.MkdirAll(path.Join(c.params.ProjectDirectory, "internal", "domain"), defaultFolderPerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	return nil
}

func (c *CommonCreator) createUsecaseFolder() error {
	err := os.MkdirAll(path.Join(c.params.ProjectDirectory, "internal", "usecase"), defaultFolderPerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	return nil
}

func (c *CommonCreator) createGitignoreFile() error {
	return createFile(c.params.ProjectDirectory, ".gitignore", gitignorePath,
		map[string]string{},
	)
}

func (c *CommonCreator) createMakefileFile() error {
	return createFile(c.params.ProjectDirectory, "Makefile", makefilePath,
		map[string]string{
			"AppName": c.params.AppName,
		},
	)
}

func (c *CommonCreator) createDockerfileFile() error {
	return createFile(c.params.ProjectDirectory, "Dockerfile", dockerfilePath,
		map[string]string{
			"AppName":   c.params.AppName,
			"GoVersion": c.params.GoVersion,
		},
	)
}

func (c *CommonCreator) createDockerComposeFile() error {
	return createFile(c.params.ProjectDirectory, "docker-compose.yml", dockerComposePath,
		map[string]string{
			"AppName": c.params.AppName,
		},
	)
}

func (c *CommonCreator) createLocalEnvFile() error {
	return createFile(c.params.ProjectDirectory, "local.env", localEnvPath,
		map[string]string{},
	)
}

func (c *CommonCreator) createProdEnvFile() error {
	return createFile(c.params.ProjectDirectory, "prod.env", prodEnvPath,
		map[string]string{},
	)
}

func (c *CommonCreator) createDockerEnvFile() error {
	return createFile(c.params.ProjectDirectory, "docker.env", dockerEnvPath,
		map[string]string{},
	)
}

func (c *CommonCreator) createLocalConfigFile() error {
	return createFile(path.Join(c.params.ProjectDirectory, "configs"), "values-local.yaml", valuesLocalPath,
		map[string]string{
			"AppName": c.params.AppName,
		},
	)
}

func (c *CommonCreator) createProdConfigFile() error {
	return createFile(path.Join(c.params.ProjectDirectory, "configs"), "values-prod.yaml", valuesProdPath,
		map[string]string{
			"AppName": c.params.AppName,
		},
	)
}

func (c *CommonCreator) createDockerConfigFile() error {
	return createFile(path.Join(c.params.ProjectDirectory, "configs"), "values-docker.yaml", valuesDockerPath,
		map[string]string{
			"AppName": c.params.AppName,
		},
	)
}

func (c *CommonCreator) createGoModFile() error {
	return createFile(c.params.ProjectDirectory, "go.mod", goModPath,
		map[string]string{
			"Module":    c.params.ModuleName,
			"GoVersion": c.params.GoVersion,
		},
	)
}

func (c *CommonCreator) createMainFile() error {
	return createFile(path.Join(c.params.ProjectDirectory, "cmd", c.params.AppName), "main.go", mainPath,
		map[string]string{
			"Module": c.params.ModuleName,
		},
	)
}

func (c *CommonCreator) createAppFile() error {
	return createFile(path.Join(c.params.ProjectDirectory, "internal", "app"), "app.go", appPath,
		map[string]string{
			"Module": c.params.ModuleName,
		},
	)
}

func (c *CommonCreator) createConfigFile() error {
	return createFile(path.Join(c.params.ProjectDirectory, "internal", "app"), "config.go", configPath,
		map[string]string{},
	)
}
