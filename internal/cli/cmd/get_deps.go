package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/DmitySH/go-service-sketch/internal/pkg/fileutils"
	"github.com/DmitySH/go-service-sketch/internal/pkg/logger"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
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

var getDepsCmd = &cobra.Command{
	Use:   "get-deps [OPTIONS]",
	Short: fmt.Sprintf("Download buf+gRPC generators to .%s folder", binDir),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		ctx, cancel := context.WithTimeout(ctx, time.Second*20)
		defer cancel()

		if err := os.MkdirAll(binDir, 0750); err != nil {
			logger.Fatalf(ctx, "can't create dir for deps: %v", err)
		}

		eg, egCtx := errgroup.WithContext(ctx)
		for fileName, url := range binDeps {
			eg.Go(func() error {
				return fileutils.DownloadFile(egCtx, filepath.Join(binDir, fileName), url)
			})
		}

		if err := eg.Wait(); err != nil {
			logger.Fatal(ctx, err.Error())
		}
	},
}
