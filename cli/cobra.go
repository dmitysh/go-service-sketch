package cli

import (
	"github.com/spf13/cobra"
)

func SetupRootCommand(rootCmd *cobra.Command) {
	rootCmd.Annotations = map[string]string{
		"additionalHelp": "For more help on how to use Sketch contact Telegram: @DmitriyShagarov",
	}
}
