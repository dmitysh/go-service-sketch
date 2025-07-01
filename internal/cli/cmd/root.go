package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)

	jojoCmd.AddCommand(getDepsCmd)
	jojoCmd.AddCommand(generateCmd)

	addCmd.AddCommand(addPostgresCmd)
	addCmd.AddCommand(addGrpcCmd)

	rootCmd.AddCommand(
		initCmd,
		jojoCmd,
		addCmd,
	)
}

var rootCmd = &cobra.Command{
	Use:   "sketch",
	Short: "Tool for fast file drop between machines in local net",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Annotations: map[string]string{
		"additionalHelp": "For more help on how to use Sketch contact Telegram: @DmitriyShagarov",
	},
}

// Execute executes the root command
func Execute() {
	_ = rootCmd.Execute()
}
