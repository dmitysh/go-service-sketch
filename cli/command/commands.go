package command

import (
	"github.com/DmitySH/go-service-sketch/cli"
	"github.com/spf13/cobra"
)

// AddAllCommands Adds all the commands from cli/command to the root command
func AddAllCommands(cmd *cobra.Command, sketchCli *cli.SketchCli) {
	jojo := NewjojoCommand(sketchCli)
	jojo.AddCommand(NewGetDepsCommand(sketchCli))
	jojo.AddCommand(NewGenerateCommand(sketchCli))

	cmd.AddCommand(
		NewInitCommand(sketchCli),
		jojo,
	)
}
