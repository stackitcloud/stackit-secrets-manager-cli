package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	deleteUserInstanceId string
)

var deleteUserCmd = &cobra.Command{
	Use:   "users <userId>",
	Short: "Deletes a user from the Secrets Manager.",
	Long:  `Deletes a user from the Secrets Manager.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, userId := range args {
			if err := deleteUser(userId); err != nil {
				fmt.Printf("ERROR: %v\n", err)
			}
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteUserCmd)

	deleteUserCmd.PersistentFlags().StringVar(&deleteUserInstanceId, "instance-id", "", "The UUID of the instance to get the users for.")
	_ = deleteUserCmd.MarkPersistentFlagRequired("instance-id")
}

func deleteUser(userId string) error {
	client, err := createClient()
	if err != nil {
		return err
	}
	response, err := client.DeleteV1ProjectsProjectIdInstancesInstanceIdUsersUserId(context.Background(), projectId, deleteUserInstanceId, userId)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}
	fmt.Printf("User %s deleted\n", userId)
	return nil
}
