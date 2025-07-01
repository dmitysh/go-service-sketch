package cmd

import (
	"github.com/DmitySH/go-service-sketch/internal/pkg/logger"
	"github.com/DmitySH/go-service-sketch/internal/temple"
	"github.com/spf13/cobra"
)

var addPostgresCmd = &cobra.Command{
	Use:   "postgres",
	Short: "Add PostgreSQL",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		creator := temple.NewDBCreator()
		err := creator.Create()
		if err != nil {
			logger.Fatalf(ctx, "can't add postgres: %v", err)
		}
		logger.Info(ctx, "Success! Use go mod tidy to complete")
	},
}
