package cmd

import (
	"github.com/DmitySH/go-service-sketch/internal/pkg/logger"
	"github.com/DmitySH/go-service-sketch/internal/temple"
	"github.com/spf13/cobra"
)

type initOptions struct {
	outDir     string
	appName    string
	moduleName string
	goVersion  string
}

var initOpts initOptions

func init() {
	const (
		moduleFlagName = "module"
		goverFlagName  = "go-ver"
	)

	flags := initCmd.Flags()

	flags.StringVarP(&initOpts.appName, "name", "n", "app", "Service name")
	flags.StringVarP(&initOpts.outDir, "out", "o", ".", "Path to directory with generated service template")
	flags.StringVar(&initOpts.moduleName, moduleFlagName, "", "Name of module in go mod init command")
	flags.StringVar(&initOpts.goVersion, goverFlagName, "", "Go version")

	_ = initCmd.MarkFlagRequired(moduleFlagName)
	_ = initCmd.MarkFlagRequired(goverFlagName)
}

var initCmd = &cobra.Command{
	Use:   "init [OPTIONS]",
	Short: "Generate service template",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		creators := initAllCreators(initOpts)

		for _, creator := range creators {
			if err := creator.Create(); err != nil {
				logger.Fatal(ctx, err.Error())
			}
		}

		logger.Info(ctx, "Success! Use go mod tidy to complete initialization")
	},
}

// TODO: fmt -> logger

type Creator interface {
	Create() error
}

func initAllCreators(opts initOptions) []Creator {
	var creators []Creator

	creators = append(creators,
		temple.NewCommonCreator(temple.CommonCreatorParams{
			ProjectDirectory: opts.outDir,
			AppName:          opts.appName,
			ModuleName:       opts.moduleName,
			GoVersion:        opts.goVersion,
		}),
		temple.NewHoustonCreator(opts.moduleName, opts.outDir),
	)

	return creators
}
