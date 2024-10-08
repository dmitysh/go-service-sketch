package command

import (
	"context"
	"fmt"
	"github.com/DmitySH/go-service-sketch/pkg/fileutils"
	"golang.org/x/sync/errgroup"
	"os"
	"path/filepath"
	"time"

	"github.com/DmitySH/go-service-sketch/cli"
	"github.com/DmitySH/go-service-sketch/cli/temple"
	"github.com/spf13/cobra"
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

func NewAddGRPCCommand(sketchCli *cli.SketchCli) *cobra.Command {
	var options addGRPCOptions

	cmd := &cobra.Command{
		Use:   "grpc",
		Short: "Add gRPC server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runAddGRPC(cmd, sketchCli, &options)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&options.serviceName, "service-name", "n", "", "Service name")

	return cmd
}

func runAddGRPC(cmd *cobra.Command, _ *cli.SketchCli, options *addGRPCOptions) error {
	creator := temple.NewGRPCCreator(options.serviceName)
	err := creator.Create()
	if err != nil {
		return fmt.Errorf("can't add grpc: %w", err)
	}

	if !fileutils.IsDirExists(swaggerFilesDir) {
		err = getSwaggerFiles(cmd.Context())
		if err != nil {
			return fmt.Errorf("can't download swagger files: %w", err)
		}
	}

	fmt.Println("Success! Use jojo command and go mod tidy to complete")
	fmt.Println("Don't forget to add server section to values-*.yaml config files!")

	return nil
}

func getSwaggerFiles(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	if err := os.MkdirAll(swaggerFilesDir, 0755); err != nil {
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
