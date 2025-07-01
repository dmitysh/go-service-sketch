package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/dmitysh/go-service-sketch/internal/pkg/fileutils"
	"github.com/dmitysh/go-service-sketch/internal/pkg/logger"
	"github.com/dmitysh/go-service-sketch/internal/temple"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

const (
	swaggerFilesDir = "internal/server/swagger"
)

var (
	swaggerFiles = map[string]string{
		"favicon-16x16.png":               "https://raw.githubusercontent.com/DmitySH/go-service-sketch/main/swagger-ui/favicon-16x16.png",
		"favicon-32x32.png":               "https://raw.githubusercontent.com/DmitySH/go-service-sketch/main/swagger-ui/favicon-32x32.png",
		"index.css":                       "https://raw.githubusercontent.com/DmitySH/go-service-sketch/main/swagger-ui/index.css",
		"index.html":                      "https://raw.githubusercontent.com/DmitySH/go-service-sketch/main/swagger-ui/index.html",
		"oauth2-redirect.html":            "https://raw.githubusercontent.com/DmitySH/go-service-sketch/main/swagger-ui/oauth2-redirect.html",
		"swagger-initializer.js":          "https://raw.githubusercontent.com/DmitySH/go-service-sketch/main/swagger-ui/swagger-initializer.js",
		"swagger-ui-bundle.js":            "https://raw.githubusercontent.com/DmitySH/go-service-sketch/main/swagger-ui/swagger-ui-bundle.js",
		"swagger-ui-es-bundle-core.js":    "https://raw.githubusercontent.com/DmitySH/go-service-sketch/main/swagger-ui/swagger-ui-es-bundle-core.js",
		"swagger-ui-es-bundle.js":         "https://raw.githubusercontent.com/DmitySH/go-service-sketch/main/swagger-ui/swagger-ui-es-bundle.js",
		"swagger-ui-standalone-preset.js": "https://raw.githubusercontent.com/DmitySH/go-service-sketch/main/swagger-ui/swagger-ui-standalone-preset.js",
		"swagger-ui.css":                  "https://raw.githubusercontent.com/DmitySH/go-service-sketch/main/swagger-ui/swagger-ui.css",
		"swagger-ui.js":                   "https://raw.githubusercontent.com/DmitySH/go-service-sketch/main/swagger-ui/swagger-ui.js",
	}
)

type addGRPCOptions struct {
	serviceName string
}

var addGRPCOpts addGRPCOptions

func init() {
	const (
		serviceNameFlagName = "service-name"
	)

	flags := addGrpcCmd.Flags()

	flags.StringVarP(&addGRPCOpts.serviceName, "service-name", "n", "", "Service name")

	_ = addGrpcCmd.MarkFlagRequired(serviceNameFlagName)
}

var addGrpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Add gRPC server",
	Run: func(cmd *cobra.Command, _ []string) {
		ctx := cmd.Context()

		creator := temple.NewGRPCCreator(addGRPCOpts.serviceName)
		err := creator.Create()
		if err != nil {
			logger.Fatalf(ctx, "can't add grpc: %v", err)
		}

		if !fileutils.IsDirExists(swaggerFilesDir) {
			err = getSwaggerFiles(cmd.Context())
			if err != nil {
				logger.Fatalf(ctx, "can't download swagger files: %v", err)
			}
		}

		logger.Info(ctx, "Success! Use jojo command and go mod tidy to complete")
		logger.Info(ctx, "Don't forget to add server section to values-*.yaml config files!")
	},
}

func getSwaggerFiles(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	if err := os.MkdirAll(swaggerFilesDir, 0750); err != nil {
		return fmt.Errorf("can't create dir for files: %w", err)
	}

	eg, egCtx := errgroup.WithContext(ctx)
	for fileName, url := range swaggerFiles {
		eg.Go(func() error {
			return fileutils.DownloadFile(egCtx, filepath.Join(swaggerFilesDir, fileName), url)
		})
	}

	return eg.Wait()
}
