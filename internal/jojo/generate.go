package jojo

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/DmitySH/go-service-sketch/internal/jojo/buf"
	"github.com/DmitySH/go-service-sketch/internal/pkg/fileutils"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v3"
)

const (
	configPath        = "jojo.yaml"
	bufBinPath        = "bin/buf"
	tmpGenerateFolder = ".generate-tmp"
	pbFolder          = "internal/pb"
	goModPath         = "go.mod"
)

func Generate(ctx context.Context, preserveGenerated bool) error {
	cfg, err := readConfig(configPath)
	if err != nil {
		return fmt.Errorf("can't read config: %w", err)
	}
	_ = cfg

	if err = os.MkdirAll(tmpGenerateFolder, 0750); err != nil {
		return fmt.Errorf("can't create tmp folder: %w", err)
	}
	defer func() {
		if preserveGenerated {
			return
		}

		if err = os.RemoveAll(tmpGenerateFolder); err != nil {
			fmt.Printf("can't remove tmp folder: %v. you should delete it manually\n", err)
		}
	}()

	protoPaths := make([]string, 0)

	for _, localDep := range cfg.LocalDependencies {
		fullFilePath := filepath.Join(tmpGenerateFolder, localDep.FilePath)

		protoPaths = append(protoPaths, fullFilePath)

		err = os.MkdirAll(filepath.Dir(fullFilePath), 0750)
		if err != nil {
			return fmt.Errorf("can't create dir for proto file: %w", err)
		}

		err = fileutils.CopyFile(localDep.Location, fullFilePath)
		if err != nil {
			return fmt.Errorf("can't copy local dependecy to tmp dir: %w", err)
		}
	}

	eg, egCtx := errgroup.WithContext(ctx)
	for _, extDep := range cfg.ExternalDependencies {
		fullFilePath := filepath.Join(tmpGenerateFolder, extDep.FilePath)

		eg.Go(func() error {
			if err = os.MkdirAll(filepath.Dir(fullFilePath), 0750); err != nil {
				return fmt.Errorf("can't create dir: %w", err)
			}

			return fileutils.DownloadFile(egCtx, fullFilePath, extDep.URL)
		})

		if !extDep.NoGenerate {
			protoPaths = append(protoPaths, fullFilePath)
		}
	}
	if err = eg.Wait(); err != nil {
		return fmt.Errorf("can't download external dependency: %w", err)
	}

	bufTmpl, err := overrideBufImportPaths(buf.DefaultBufGenConfig(), protoPaths)
	if err != nil {
		return fmt.Errorf("can't override buf imports: %w", err)
	}

	if err = createBufTemplate(bufTmpl); err != nil {
		return fmt.Errorf("can't create template: %w", err)
	}

	genBackend := buf.New(tmpGenerateFolder, bufBinPath, protoPaths)
	err = genBackend.Execute(ctx)
	if err != nil {
		return fmt.Errorf("can't execute generation command: %w", err)
	}

	if err = fileutils.CopyDir(filepath.Join(tmpGenerateFolder, buf.GeneratedFolder), pbFolder); err != nil {
		return fmt.Errorf("can't copy generated files: %w", err)
	}

	return nil
}

func createBufTemplate(t buf.BufGenConfig) error {
	f, err := os.Create(filepath.Join(tmpGenerateFolder, buf.TemplateFile))
	if err != nil {
		return fmt.Errorf("can't create template file: %w", err)
	}
	defer f.Close()

	enc := yaml.NewEncoder(f)
	err = enc.Encode(t)
	if err != nil {
		return fmt.Errorf("can't encode buf config: %w", err)
	}

	err = enc.Close()
	if err != nil {
		return fmt.Errorf("can't close yaml encoder: %w", err)
	}

	return nil
}

func overrideBufImportPaths(t buf.BufGenConfig, protoPaths []string) (buf.BufGenConfig, error) {
	module, err := getProjectGoModule()
	if err != nil {
		return buf.BufGenConfig{}, fmt.Errorf("can't get go module: %w", err)
	}

	for _, path := range protoPaths {
		cleanPath := strings.ReplaceAll(path, tmpGenerateFolder+string(filepath.Separator), "")
		for i := 0; i < len(t.Plugins); i++ {
			t.Plugins[i].Opt = append(t.Plugins[i].Opt, fmt.Sprintf("M%s=%s", cleanPath, filepath.Join(module, pbFolder, filepath.Dir(cleanPath))))
		}
	}

	return t, nil
}

func getProjectGoModule() (string, error) {
	f, err := os.Open(goModPath)
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
