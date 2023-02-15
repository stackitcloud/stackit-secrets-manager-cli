package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-secrets-manager-cli/internal/api"
	"net/http"
)

var (
	updateInstanceinstanceId   string
	updateInstanceUserLimit    int
	updateInstanceSecretLimit  int
	updateInstanceVersionLimit int
)

var updateInstanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "Updates an instance for the Secrets Manager.",
	Long:  `Updates an instance for the Secrets Manager.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := updateInstance(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	},
}

func init() {
	updateCmd.AddCommand(updateInstanceCmd)

	updateInstanceCmd.PersistentFlags().StringVar(&updateInstanceinstanceId, "instance-id", "", "The UUID of the instance to update.")
	_ = updateInstanceCmd.MarkPersistentFlagRequired("instance-id")
	updateInstanceCmd.PersistentFlags().IntVar(&updateInstanceUserLimit, "user-limit", 5, "The number of maximum users, 5-100 in steps of 5.")
	updateInstanceCmd.PersistentFlags().IntVar(&updateInstanceSecretLimit, "secret-limit", 100, "The number of maximum secrets, 100-1.000 in steps of 100.")
	updateInstanceCmd.PersistentFlags().IntVar(&updateInstanceVersionLimit, "version-limit", 0, "The number of maximum versions, 0 or 5.")
}

func updateInstance() error {
	client, err := createClient()
	if err != nil {
		return err
	}
	request := api.PutV1ProjectsProjectIdInstancesInstanceIdJSONRequestBody{
		UserLimit:    updateInstanceUserLimit,
		SecretLimit:  updateInstanceSecretLimit,
		VersionLimit: updateInstanceVersionLimit,
	}
	response, err := client.PutV1ProjectsProjectIdInstancesInstanceId(context.Background(), projectId, updateInstanceinstanceId, request)
	if err != nil {
		return fmt.Errorf("failed to update instance: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}
	return nil
}
