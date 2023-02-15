package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes resources from the Secrets Manager.",
	Long:  `Deletes resources from the Secrets Manager.`,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
