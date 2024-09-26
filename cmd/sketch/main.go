package main

import (
	"os"

	"github.com/DmitySH/go-service-sketch/cli"
	"github.com/DmitySH/go-service-sketch/cli/command"
	"github.com/spf13/cobra"
)

const version = "0.0.1"

func main() {
	sketchCli := cli.NewSketchCli()

	if runSketchErr := runSketch(sketchCli); runSketchErr != nil {
		os.Exit(1)
	}
}

func runSketch(sketchCli *cli.SketchCli) error {
	topCmd := newSketchCommand(sketchCli)
	return topCmd.Execute()
}

func newSketchCommand(sketchCli *cli.SketchCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sketch",
		Short:   "Tool for creating Go services",
		Version: version,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	cli.SetupRootCommand(cmd)
	command.AddAllCommands(cmd, sketchCli)

	return cmd
}
