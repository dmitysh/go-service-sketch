package command

import (
	"github.com/DmitySH/go-service-sketch/cli"
	"github.com/spf13/cobra"
)

func NewjojoCommand(_ *cli.SketchCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "jojo [COMMAND]",
		Short: "Command to make life with Protobufs easier",
	}

	return cmd
}
