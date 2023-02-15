package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates, updates and deletes resources for the Secrets Manager.",
	Long:  `Creates, updates and deletes resources for the Secrets Manager.`,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
