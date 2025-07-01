package buf

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	GeneratedFolder = ".generated"
)

type Buf struct {
	workingDir string
	binPath    string
	protoPaths []string
}

func New(workingDir string, binPath string, protoPaths []string) Buf {
	return Buf{
		workingDir: workingDir,
		binPath:    binPath,
		protoPaths: protoPaths,
	}
}

func (b Buf) Execute(ctx context.Context) error {
	cmdArgs := make([]string, 0)
	cmdArgs = append(cmdArgs, "generate")
	cmdArgs = append(cmdArgs, "--template")
	cmdArgs = append(cmdArgs, filepath.Join(b.workingDir, TemplateFile))
	cmdArgs = append(cmdArgs, "--output")
	cmdArgs = append(cmdArgs, filepath.Join(b.workingDir, GeneratedFolder))

	for _, path := range b.protoPaths {
		cmdArgs = append(cmdArgs, "--path")
		cmdArgs = append(cmdArgs, path)
	}

	cmdArgs = append(cmdArgs, b.workingDir)

	cmd := exec.CommandContext(ctx, b.binPath, cmdArgs...) //nolint:gosec
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
