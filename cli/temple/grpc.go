package temple

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type GRPCCreator struct {
	serviceName    string
	folderCreators []func() error
	fileCreators   []func() error
}

func NewGRPCCreator(serviceName string) *GRPCCreator {
	g := &GRPCCreator{
		serviceName: serviceName,
	}

	g.folderCreators = []func() error{
		g.createServerFolder,
	}

	g.fileCreators = []func() error{
		g.createProtoFile,
		g.createJojoFile,
		g.createGrutResponseFile,
		g.createGrutInterceptorFile,
		g.createServerGRPCFile,
		g.createServerHTTPFile,
		g.createServerConfigFile,
	}

	return g
}

func (g *GRPCCreator) Create() error {
	if err := createFiles(g.fileCreators); err != nil {
		return err
	}

	if err := createFolders(g.fileCreators); err != nil {
		return err
	}

	return nil
}

func (g *GRPCCreator) createServerFolder() error {
	err := os.MkdirAll(path.Join("internal", "server"), defaultFolderPerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	return nil
}

func (g *GRPCCreator) createProtoFile() error {
	serviceName, err := g.getServiceName()
	if err != nil {
		return fmt.Errorf("can't get service name: %w", err)
	}

	module, err := getProjectGoModule()
	if err != nil {
		return fmt.Errorf("can't get go module: %w", err)
	}

	return createFile(filepath.Join("api", serviceName), "service.proto", protoPath,
		map[string]string{
			"Module":           filepath.Join(module, "api", serviceName),
			"ServiceNameCamel": kebabToCamel(serviceName),
			"ServiceNameSnake": kebabToSnake(serviceName),
		},
	)
}

func (g *GRPCCreator) createJojoFile() error {
	serviceName, err := g.getServiceName()
	if err != nil {
		return fmt.Errorf("can't get service name: %w", err)
	}

	return createFile(".", "jojo.yaml", jojoPath,
		map[string]string{
			"ServiceName": serviceName,
		},
	)
}

func (g *GRPCCreator) createGrutInterceptorFile() error {
	module, err := getProjectGoModule()
	if err != nil {
		return fmt.Errorf("can't get go module: %w", err)
	}

	return createFile(houstonPath("grut"), "interceptor.go", houstonGrutInterceptorPath,
		map[string]string{
			"Module": module,
		},
	)
}

func (g *GRPCCreator) createGrutResponseFile() error {
	module, err := getProjectGoModule()
	if err != nil {
		return fmt.Errorf("can't get go module: %w", err)
	}

	return createFile(houstonPath("grut"), "response.go", houstonGrutResponsePath,
		map[string]string{
			"Module": module,
		},
	)
}

func (g *GRPCCreator) createServerGRPCFile() error {
	serviceName, err := g.getServiceName()
	if err != nil {
		return fmt.Errorf("can't get service name: %w", err)
	}

	module, err := getProjectGoModule()
	if err != nil {
		return fmt.Errorf("can't get go module: %w", err)
	}

	return createFile(filepath.Join("internal", "server"), "grpc.go", serverGRPCPath,
		map[string]string{
			"Module":           module,
			"ServiceName":      serviceName,
			"ServiceNameCamel": kebabToCamel(serviceName),
			"ServiceNameSnake": kebabToSnake(serviceName),
		},
	)
}

func (g *GRPCCreator) createServerHTTPFile() error {
	serviceName, err := g.getServiceName()
	if err != nil {
		return fmt.Errorf("can't get service name: %w", err)
	}

	module, err := getProjectGoModule()
	if err != nil {
		return fmt.Errorf("can't get go module: %w", err)
	}

	return createFile(filepath.Join("internal", "server"), "http.go", serverHTTPPath,
		map[string]string{
			"Module":           module,
			"ServiceName":      serviceName,
			"ServiceNameCamel": kebabToCamel(serviceName),
			"ServiceNameSnake": kebabToSnake(serviceName),
		},
	)
}

func (g *GRPCCreator) createServerConfigFile() error {
	return createFile(filepath.Join("internal", "server"), "config.go", serverConfigPath,
		map[string]string{},
	)
}

func (g *GRPCCreator) getServiceName() (string, error) {
	var serviceName string
	if g.serviceName != "" {
		serviceName = g.serviceName
	} else {
		projectDir, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("can't get project dir: %w", err)
		}
		serviceName = filepath.Base(projectDir)
	}

	return serviceName, nil
}

func kebabToCamel(kebab string) string {
	words := strings.Split(kebab, "-")

	if len(words) == 0 {
		return ""
	}

	var camelCase string
	for _, word := range words {
		camelCase += strings.Title(word)
	}

	return camelCase
}

func kebabToSnake(kebab string) string {
	return strings.ReplaceAll(kebab, "-", "_")
}
