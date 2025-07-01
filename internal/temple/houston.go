package temple

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
		h.createSecretGetFile,
		h.createSecretInitFile,
		h.createStageFile,
	}

	return h
}

func (h *HoustonCreator) Create() error {
	if err := createFiles(h.fileCreators); err != nil {
		return err
	}

	return nil
}

func (h *HoustonCreator) createCloserFile() error {
	return createFile(houstonPath("closer"), "closer.go", houstonCloserPath,
		map[string]string{
			"Module": h.module,
		},
	)
}

func (h *HoustonCreator) createConfigGetFile() error {
	return createFile(houstonPath("config"), "get.go", houstonConfigGetPath,
		map[string]string{
			"Module": h.module,
		},
	)
}

func (h *HoustonCreator) createConfigYAMLFile() error {
	return createFile(houstonPath("config"), "yaml.go", houstonConfigYAMLPath,
		map[string]string{
			"Module": h.module,
		},
	)
}

func (h *HoustonCreator) createLoggyInitFile() error {
	return createFile(houstonPath("loggy"), "init.go", houstonLoggyInitPath,
		map[string]string{},
	)
}

func (h *HoustonCreator) createLoggyLogFile() error {
	return createFile(houstonPath("loggy"), "log.go", houstonLoggyLogPath,
		map[string]string{},
	)
}

func (h *HoustonCreator) createSecretEnvFile() error {
	return createFile(houstonPath("secret"), "env.go", houstonSecretEnvPath,
		map[string]string{},
	)
}

func (h *HoustonCreator) createSecretGetFile() error {
	return createFile(houstonPath("secret"), "get.go", houstonSecretGetPath,
		map[string]string{},
	)
}

func (h *HoustonCreator) createSecretInitFile() error {
	return createFile(houstonPath("secret"), "init.go", houstonSecretInitPath,
		map[string]string{},
	)
}

func (h *HoustonCreator) createStageFile() error {
	return createFile(houstonPath("stage"), "stage.go", houstonStagePath,
		map[string]string{},
	)
}
