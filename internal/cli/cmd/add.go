package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [COMMAND]",
	Short: "Adds resource to existing sketch project",
}
