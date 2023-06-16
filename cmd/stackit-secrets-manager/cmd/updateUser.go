package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-secrets-manager-cli/internal/api"
)

var (
	updateUserInstanceId  string
	updateUserUserId      string
	updateUserEnableWrite bool
)

var updateUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Updates an user for the Secrets Manager.",
	Long:  `Updates an user for the Secrets Manager.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := updateUser(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	},
}

func init() {
	updateCmd.AddCommand(updateUserCmd)

	updateUserCmd.PersistentFlags().StringVar(&updateUserInstanceId, "instance-id", "", "The UUID of the instance to update the user for.")
	_ = updateUserCmd.MarkPersistentFlagRequired("instance-id")
	updateUserCmd.PersistentFlags().StringVar(&updateUserUserId, "user-id", "", "The UUID of the user to update.")
	_ = updateUserCmd.MarkPersistentFlagRequired("user-id")
	updateUserCmd.PersistentFlags().BoolVar(&updateUserEnableWrite, "enable-write", false, "Update write permissions for the user.")
}

func updateUser() error {
	client, err := createClient()
	if err != nil {
		return err
	}
	request := api.PutV1ProjectsProjectIdInstancesInstanceIdUsersUserIdJSONRequestBody{
		Write: &updateUserEnableWrite,
	}
	response, err := client.PutV1ProjectsProjectIdInstancesInstanceIdUsersUserId(context.Background(), projectId, updateUserInstanceId, updateUserUserId, request)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}
	fmt.Printf("User %s updated\n", updateUserUserId)
	return nil
}
