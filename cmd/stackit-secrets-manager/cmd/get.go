package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns resources from the Secrets Manager.",
	Long:  `Returns resources from the Secrets Manager.`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
