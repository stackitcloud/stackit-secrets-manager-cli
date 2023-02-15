package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stackitcloud/stackit-secrets-manager-cli/internal/api"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	createUserInstanceId  string
	createUserDescription string
	createUserEnableWrite bool
)

var createUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Creates a new user for the Secrets Manager.",
	Long:  `Creates a new user for the Secrets Manager.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := createUser(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	},
}

func init() {
	createCmd.AddCommand(createUserCmd)

	createUserCmd.PersistentFlags().StringVar(&createUserInstanceId, "instance-id", "", "The UUID of the instance to create the user for.")
	_ = createUserCmd.MarkPersistentFlagRequired("instance-id")
	createUserCmd.PersistentFlags().StringVar(&createUserDescription, "description", "", "The description to associate the new user with.")
	createUserCmd.PersistentFlags().BoolVar(&createUserEnableWrite, "enable-write", false, "Enable write permissions for the user.")
}

func createUser() error {
	client, err := createClient()
	if err != nil {
		return err
	}
	request := api.PostV1ProjectsProjectIdInstancesInstanceIdUsersJSONRequestBody{
		Description: createUserDescription,
		Write:       createUserEnableWrite,
	}
	response, err := client.PostV1ProjectsProjectIdInstancesInstanceIdUsers(context.Background(), projectId, createUserInstanceId, request)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}

	var user api.User
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}
	if err := printUsersWithPassword([]api.User{user}); err != nil {
		return err
	}
	return nil
}
