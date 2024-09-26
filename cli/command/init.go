package command

import (
	"fmt"
	"github.com/DmitySH/go-service-sketch/cli"
	"github.com/DmitySH/go-service-sketch/cli/temple"
	"github.com/spf13/cobra"
)

type Creator interface {
	Create() error
}

type initOptions struct {
	outDir     string
	appName    string
	moduleName string
	goVersion  string
	grpc       bool
}

func NewInitCommand(sketchCli *cli.SketchCli) *cobra.Command {
	var options initOptions

	cmd := &cobra.Command{
		Use:   "init [OPTIONS]",
		Short: "Generate service template",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInit(cmd, sketchCli, &options)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&options.appName, "name", "n", "app", "Service name")
	flags.StringVarP(&options.outDir, "out", "o", ".", "Path to directory with generated service template")
	flags.StringVar(&options.moduleName, "module", "", "Name of module in go mod init command")
	flags.StringVar(&options.goVersion, "go-ver", "", "Go version")
	//flags.BoolVar(&options.grpc, "grpc", false, "Generate gRPC server")

	_ = cmd.MarkFlagRequired("module")
	_ = cmd.MarkFlagRequired("go-ver")

	return cmd
}

func runInit(_ *cobra.Command, _ *cli.SketchCli, options *initOptions) error {
	creators := initAllCreators(options)

	for _, creator := range creators {
		if err := creator.Create(); err != nil {
			return err
		}
	}

	fmt.Println("Success! Use go mod tidy to complete initialization")

	return nil
}

func initAllCreators(opts *initOptions) []Creator {
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
