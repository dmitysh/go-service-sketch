package temple

import (
	"bufio"
	"embed"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

const (
	defaultFolderPerm os.FileMode = 0755
)

//go:embed templates
var templatesFolder embed.FS

// Common
const (
	makefilePath      = "templates/makefile.txt"
	gitignorePath     = "templates/gitignore.txt"
	dockerfilePath    = "templates/dockerfile.txt"
	dockerComposePath = "templates/docker-compose.txt"
	localEnvPath      = "templates/local.txt"
	prodEnvPath       = "templates/prod.txt"
	dockerEnvPath     = "templates/docker.txt"
	valuesLocalPath   = "templates/configs/values-local.txt"
	valuesProdPath    = "templates/configs/values-prod.txt"
	valuesDockerPath  = "templates/configs/values-docker.txt"
	goModPath         = "templates/go.mod.txt"
	mainPath          = "templates/cmd/app/main.txt"
	appPath           = "templates/internal/app/app.txt"
	configPath        = "templates/internal/app/config.txt"
)

// Houston
const (
	houstonCloserPath           = "templates/internal/pkg/houston/closer/closer.txt"
	houstonConfigGetPath        = "templates/internal/pkg/houston/config/get.txt"
	houstonConfigYAMLPath       = "templates/internal/pkg/houston/config/yaml.txt"
	houstonLoggyInitPath        = "templates/internal/pkg/houston/loggy/init.txt"
	houstonLoggyLogPath         = "templates/internal/pkg/houston/loggy/log.txt"
	houstonSecretEnvPath        = "templates/internal/pkg/houston/secret/env.txt"
	houstonSecretGetPath        = "templates/internal/pkg/houston/secret/get.txt"
	houstonSecretInitPath       = "templates/internal/pkg/houston/secret/init.txt"
	houstonStagePath            = "templates/internal/pkg/houston/stage/stage.txt"
	houstonDobbyPGXPath         = "templates/internal/pkg/houston/dobby/pgx.txt"
	houstonDobbyTransactionPath = "templates/internal/pkg/houston/dobby/transaction.txt"
	houstonGrutInterceptorPath  = "templates/internal/pkg/houston/grut/interceptor.txt"
	houstonGrutResponsePath     = "templates/internal/pkg/houston/grut/response.txt"
)

// gRPC
const (
	protoPath        = "templates/api/app/service.txt"
	jojoPath         = "templates/jojo.txt"
	serverGRPCPath   = "templates/internal/server/grpc.txt"
	serverHTTPPath   = "templates/internal/server/http.txt"
	serverConfigPath = "templates/internal/server/config.txt"
)

// Other
const (
	projectGoModPath = "go.mod"
)

func houstonPath(lastElem string) string {
	return path.Join("internal", "pkg", "houston", lastElem)
}

func getProjectGoModule() (string, error) {
	f, err := os.Open(projectGoModPath)
	if err != nil {
		return "", fmt.Errorf("can't open go mod file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	firstLine := ""
	if scanner.Scan() {
		firstLine = scanner.Text()
	} else {
		return "", errors.New("can't to read go module line")
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("scanner error: %w", err)
	}

	splittedFirstLine := strings.Split(firstLine, " ")
	if len(splittedFirstLine) != 2 {
		return "", errors.New("invalid go mod file")
	}

	return splittedFirstLine[1], nil
}
