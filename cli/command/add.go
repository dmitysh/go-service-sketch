package command

import (
	"github.com/DmitySH/go-service-sketch/cli"
	"github.com/spf13/cobra"
)

func NewAddCommand(_ *cli.SketchCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [COMMAND]",
		Short: "Adds resource to existing sketch project",
	}

	return cmd
}
