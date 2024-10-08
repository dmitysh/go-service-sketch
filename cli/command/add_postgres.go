package command

import (
	"fmt"

	"github.com/DmitySH/go-service-sketch/cli"
	"github.com/DmitySH/go-service-sketch/cli/temple"
	"github.com/spf13/cobra"
)

type addPostgresOptions struct{}

func NewAddPostgresCommand(sketchCli *cli.SketchCli) *cobra.Command {
	var options addPostgresOptions

	cmd := &cobra.Command{
		Use:   "postgres",
		Short: "Add PostgreSQL",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runAddPostgres(cmd, sketchCli, &options)
		},
	}

	return cmd
}

func runAddPostgres(_ *cobra.Command, _ *cli.SketchCli, _ *addPostgresOptions) error {
	creator := temple.NewDBCreator()
	err := creator.Create()
	if err != nil {
		return fmt.Errorf("can't add postgres: %w", err)
	}
	fmt.Println("Success! Use go mod tidy to complete")

	return nil
}
