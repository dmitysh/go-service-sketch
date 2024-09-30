package command

import (
	"context"
	"fmt"
	"github.com/DmitySH/go-service-sketch/cli"
	"github.com/DmitySH/go-service-sketch/pkg/fileutils"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
	"os"
	"path/filepath"
	"time"
)

const (
	binDir = "bin"
)

var (
	binDeps = map[string]string{
		"buf":                     "https://raw.githubusercontent.com/DmitySH/go-service-sketch/main/bin-deps/buf",
		"protoc-gen-go":           "https://raw.githubusercontent.com/DmitySH/go-service-sketch/main/bin-deps/protoc-gen-go",
		"protoc-gen-go-grpc":      "https://raw.githubusercontent.com/DmitySH/go-service-sketch/main/bin-deps/protoc-gen-go-grpc",
		"protoc-gen-grpc-gateway": "https://raw.githubusercontent.com/DmitySH/go-service-sketch/main/bin-deps/protoc-gen-grpc-gateway",
		"protoc-gen-openapiv2":    "https://raw.githubusercontent.com/DmitySH/go-service-sketch/main/bin-deps/protoc-gen-openapiv2",
	}
)

type getDepsOptions struct{}

func NewGetDepsCommand(sketchCli *cli.SketchCli) *cobra.Command {
	var options getDepsOptions

	cmd := &cobra.Command{
		Use:   "get-deps [OPTIONS]",
		Short: fmt.Sprintf("Download buf+gRPC generators to .%s folder", binDir),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGetDeps(cmd, sketchCli, &options)
		},
	}

	return cmd
}

func runGetDeps(cmd *cobra.Command, _ *cli.SketchCli, _ *getDepsOptions) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), time.Second*20)
	defer cancel()

	if err := os.MkdirAll(binDir, 0755); err != nil {
		return fmt.Errorf("can't create dir for deps: %w", err)
	}

	eg, egCtx := errgroup.WithContext(ctx)
	for fileName, url := range binDeps {
		eg.Go(func() error {
			return fileutils.DownloadFile(egCtx, filepath.Join(binDir, fileName), url)
		})
	}

	return eg.Wait()
}
