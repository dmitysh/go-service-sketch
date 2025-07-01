package jojo

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type LocalDependency struct {
	FilePath   string `yaml:"file_path"`
	Location   string `yaml:"location"`
	NoGenerate bool   `yaml:"no_generate" default:"false"`
}

type ExternalDependency struct {
	FilePath   string `yaml:"file_path"`
	URL        string `yaml:"url"`
	NoGenerate bool   `yaml:"no_generate" default:"false"`
}

type Config struct {
	Version              int                  `yaml:"version"`
	LocalDependencies    []LocalDependency    `yaml:"local_dependencies"`
	ExternalDependencies []ExternalDependency `yaml:"external_dependencies"`
}

func readConfig(fp string) (Config, error) {
	f, err := os.Open(fp) //nolint:gosec
	if err != nil {
		return Config{}, fmt.Errorf("can't open file: %w", err)
	}
	defer f.Close()

	var cfg Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("can't parse config: %w", err)
	}

	return cfg, nil
}
