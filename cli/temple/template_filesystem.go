package temple

import (
	"embed"
	"os"
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
	houstonCloserPath     = "templates/internal/pkg/houston/closer/closer.txt"
	houstonConfigGetPath  = "templates/internal/pkg/houston/config/get.txt"
	houstonConfigYAMLPath = "templates/internal/pkg/houston/config/yaml.txt"
	houstonLoggyInitPath  = "templates/internal/pkg/houston/loggy/init.txt"
	houstonLoggyLogPath   = "templates/internal/pkg/houston/loggy/log.txt"
	houstonSecretEnvPath  = "templates/internal/pkg/houston/secret/env.txt"
	houstonSecretGetPath  = "templates/internal/pkg/houston/secret/get.txt"
	houstonSecretInitPath = "templates/internal/pkg/houston/secret/init.txt"
	houstonStagePath      = "templates/internal/pkg/houston/stage/stage.txt"
)
