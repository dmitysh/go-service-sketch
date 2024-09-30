package command

import (
	"context"
	"fmt"
	"github.com/DmitySH/go-service-sketch/cli"
	"github.com/DmitySH/go-service-sketch/cli/jojo"
	"github.com/spf13/cobra"
	"time"
)

type generateOptions struct {
	preserveGenerated bool
}

func NewGenerateCommand(sketchCli *cli.SketchCli) *cobra.Command {
	var options generateOptions

	cmd := &cobra.Command{
		Use:   "generate [OPTIONS]",
		Short: "Generates pb files according to jojo.yaml",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGenerate(cmd, sketchCli, &options)
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&options.preserveGenerated, "preserve-generated", false, "Do not delete generated folder for debugging")

	return cmd
}

func runGenerate(cmd *cobra.Command, _ *cli.SketchCli, options *generateOptions) error {
	ctx, cancel := context.WithTimeout(cmd.Context(), time.Second*20)
	defer cancel()

	if err := jojo.Generate(ctx, options.preserveGenerated); err != nil {
		return err
	}

	fmt.Println("jojo magic done! Run go mod tidy to complete generation")

	return nil
}
