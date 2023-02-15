package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var deleteInstanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "Deletes an instance from the Secrets Manager.",
	Long:  `Deletes an instance from the Secrets Manager.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, instanceId := range args {
			if err := deleteInstance(instanceId); err != nil {
				fmt.Printf("ERROR: %v\n", err)
			}
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteInstanceCmd)
}

func deleteInstance(instanceId string) error {
	client, err := createClient()
	if err != nil {
		return err
	}
	response, err := client.DeleteV1ProjectsProjectIdInstancesInstanceId(context.Background(), projectId, instanceId)
	if err != nil {
		return fmt.Errorf("failed to delete instance: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}
	fmt.Printf("Instance %s deleted\n", instanceId)
	return nil
}
