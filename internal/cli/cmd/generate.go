package cmd

import (
	"context"
	"time"

	"github.com/DmitySH/go-service-sketch/internal/jojo"
	"github.com/DmitySH/go-service-sketch/internal/pkg/logger"
	"github.com/spf13/cobra"
)

type generateOptions struct {
	preserveGenerated bool
}

var generateOpts generateOptions

func init() {
	flags := generateCmd.Flags()
	flags.BoolVar(&generateOpts.preserveGenerated, "preserve-generated", false, "Do not delete generated folder for debugging")
}

var generateCmd = &cobra.Command{
	Use:   "generate [OPTIONS]",
	Short: "Generates pb files according to jojo.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		ctx, cancel := context.WithTimeout(ctx, time.Second*20)
		defer cancel()

		if err := jojo.Generate(ctx, generateOpts.preserveGenerated); err != nil {
			logger.Fatal(ctx, err.Error())
		}

		logger.Info(ctx, "jojo magic done! Run go mod tidy to complete generation")
	},
}
