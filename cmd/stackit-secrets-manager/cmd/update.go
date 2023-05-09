package cmd

import (
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates resources for the Secrets Manager.",
	Long:  `Updates resources for the Secrets Manager.`,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
