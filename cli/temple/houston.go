package temple

import (
	"go.uber.org/multierr"
	"path"
)

type HoustonCreator struct {
	module       string
	projectDir   string
	fileCreators []func() error
}

func NewHoustonCreator(module, projectDir string) *HoustonCreator {
	h := &HoustonCreator{
		module:     module,
		projectDir: projectDir,
	}

	h.fileCreators = []func() error{
		h.createCloserFile,
		h.createConfigGetFile,
		h.createConfigYAMLFile,
		h.createLoggyInitFile,
		h.createLoggyLogFile,
		h.createSecretEnvFile,
		h.createSecretGetLFile,
		h.createSecretInitFile,
		h.createStageFile,
	}

	return h
}

func (h *HoustonCreator) Create() error {
	if err := h.createFiles(); err != nil {
		return err
	}

	return nil
}

func (h *HoustonCreator) createFiles() error {
	var err error

	for _, createFn := range h.fileCreators {
		err = multierr.Append(err, createFn())
	}

	return err
}

func (h *HoustonCreator) createCloserFile() error {
	return createFile(h.houstonPath("closer"), "closer.go", houstonCloserPath,
		map[string]string{
			"Module": h.module,
		},
	)
}

func (h *HoustonCreator) createConfigGetFile() error {
	return createFile(h.houstonPath("config"), "get.go", houstonConfigGetPath,
		map[string]string{
			"Module": h.module,
		},
	)
}

func (h *HoustonCreator) createConfigYAMLFile() error {
	return createFile(h.houstonPath("config"), "yaml.go", houstonConfigYAMLPath,
		map[string]string{
			"Module": h.module,
		},
	)
}

func (h *HoustonCreator) createLoggyInitFile() error {
	return createFile(h.houstonPath("loggy"), "init.go", houstonLoggyInitPath,
		map[string]string{},
	)
}

func (h *HoustonCreator) createLoggyLogFile() error {
	return createFile(h.houstonPath("loggy"), "log.go", houstonLoggyLogPath,
		map[string]string{},
	)
}

func (h *HoustonCreator) createSecretEnvFile() error {
	return createFile(h.houstonPath("secret"), "get.go", houstonSecretGetPath,
		map[string]string{},
	)
}

func (h *HoustonCreator) createSecretGetLFile() error {
	return createFile(h.houstonPath("secret"), "get.go", houstonSecretEnvPath,
		map[string]string{},
	)
}

func (h *HoustonCreator) createSecretInitFile() error {
	return createFile(h.houstonPath("secret"), "init.go", houstonSecretInitPath,
		map[string]string{},
	)
}

func (h *HoustonCreator) createStageFile() error {
	return createFile(h.houstonPath("stage"), "stage.go", houstonStagePath,
		map[string]string{},
	)
}

func (h *HoustonCreator) houstonPath(lastElem string) string {
	return path.Join(h.projectDir, "internal", "pkg", "houston", lastElem)
}
