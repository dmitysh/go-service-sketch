package templator

import (
	"fmt"
	"github.com/DmitySH/go-service-sketch/cli"
	"github.com/spf13/cobra"
)

type kekOptions struct {
	detach     bool
	detachKeys string
}

func NewKekCommand(_ *cli.SketchCli) *cobra.Command {
	var options kekOptions

	cmd := &cobra.Command{
		Use:   "kek [OPTIONS] lol [COMMAND] [ARG...]",
		Short: "Create and run a new container from an image",
		Run:   runKek,
	}

	flags := cmd.Flags()

	flags.BoolVarP(&options.detach, "detach", "d", false, "Run container in background and print container ID")
	flags.StringVar(&options.detachKeys, "detach-keys", "", "Override the key sequence for detaching a container")

	return cmd
}

func runKek(cmd *cobra.Command, args []string) {
	fmt.Println("hello")
}
