package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates resources for the Secrets Manager.",
	Long:  `Creates resources for the Secrets Manager.`,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
