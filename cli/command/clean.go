package command

import (
	"errors"
	"github.com/DmitySH/go-service-sketch/cli"
	"github.com/DmitySH/go-service-sketch/cli/creator/config"
	"github.com/DmitySH/go-service-sketch/cli/creator/controller"
	"github.com/DmitySH/go-service-sketch/pkg/maputils"
	"github.com/spf13/cobra"
)

type Creator interface {
	Create() error
}

type cleanOptions struct {
	controller string
	outDir     string
	config     string
}

var cleanOptionsValidations = []func(options *cleanOptions) error{
	validateControllerFlag,
	validateConfigFlag,
}

var allowedControllers = map[string]struct{}{
	"grpc": {},
	"http": {},
}

var allowedConfigs = map[string]struct{}{
	"env":  {},
	"yaml": {},
}

func NewCleanCommand(sketchCli *cli.SketchCli) *cobra.Command {
	var options cleanOptions

	cmd := &cobra.Command{
		Use:   "clean [OPTIONS]",
		Short: "Generate service template using onion (clean) architecture",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validateCleanOptions(&options); err != nil {
				return err
			}

			return runClean(cmd, sketchCli, &options)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&options.outDir, "out-dir", "o", "./app", "Path to directory with generated service template")
	flags.StringVar(&options.config, "config", "", "Config type "+maputils.MapKeysToString(allowedConfigs))
	flags.StringVar(&options.controller, "controller", "", "Type of service controller "+maputils.MapKeysToString(allowedControllers))

	return cmd
}

func validateCleanOptions(options *cleanOptions) error {
	for _, validation := range cleanOptionsValidations {
		if err := validation(options); err != nil {
			return err
		}
	}

	return nil
}

func validateControllerFlag(options *cleanOptions) error {
	if _, ok := allowedControllers[options.controller]; !ok {
		return errors.New("controller must be one of " + maputils.MapKeysToString(allowedControllers))
	}

	return nil
}

func validateConfigFlag(options *cleanOptions) error {
	if _, ok := allowedConfigs[options.config]; !ok {
		return errors.New("config must be one of " + maputils.MapKeysToString(allowedConfigs))
	}

	return nil
}

func runClean(_ *cobra.Command, _ *cli.SketchCli, options *cleanOptions) error {
	creators := initAllCreators(options)

	for _, creator := range creators {
		if createErr := creator.Create(); createErr != nil {
			return createErr
		}
	}

	return nil
}

func initAllCreators(options *cleanOptions) []Creator {
	var creators []Creator

	creators = append(creators, controllerCreator(options.controller, options.outDir))
	creators = append(creators, configCreator(options.config, options.outDir))

	return creators
}

func controllerCreator(controllerName, projectDirectory string) Creator {
	switch controllerName {
	case "grpc":
		return controller.NewGRPCControllerCreator(projectDirectory)
	default:
		panic("unknown controller type")
	}
}

func configCreator(configName, projectDirectory string) Creator {
	switch configName {
	case "env":
		return config.NewEnvConfigCreator(projectDirectory)
	default:
		panic("unknown config type")
	}
}
