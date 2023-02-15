package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stackitcloud/stackit-secrets-manager-cli/internal/api"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var (
	getUsersInstanceId string
)

var getUsersCmd = &cobra.Command{
	Use:   "users",
	Short: "Returns all users from the Secrets Manager.",
	Long:  `Returns all users from the Secrets Manager.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := getUsers(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	},
}

func init() {
	getCmd.AddCommand(getUsersCmd)

	getUsersCmd.PersistentFlags().StringVar(&getUsersInstanceId, "instance-id", "", "The UUID of the instance to get the users for.")
	_ = getUsersCmd.MarkPersistentFlagRequired("instance-id")
}

func getUsers() error {
	client, err := createClient()
	if err != nil {
		return err
	}
	response, err := client.GetV1ProjectsProjectIdInstancesInstanceIdUsers(context.Background(), projectId, getUsersInstanceId)
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}

	var users api.UserList
	if err := json.NewDecoder(response.Body).Decode(&users); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}
	if err := printUsers(users.Users); err != nil {
		return err
	}
	return nil
}

func printUsers(users []api.User) error {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer writer.Flush()
	_, err := fmt.Fprintf(writer, "Id\tName\tDescription\tWrite\n")
	if err != nil {
		return fmt.Errorf("failed to write to tabwriter: %w", err)
	}
	for _, user := range users {
		_, err = fmt.Fprintf(writer, "%s\t%s\t%s\t%t\n", user.Id, user.Username, user.Description, user.Write)
		if err != nil {
			return fmt.Errorf("failed to write to tabwriter: %w", err)
		}
	}
	return nil
}

func printUsersWithPassword(users []api.User) error {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer writer.Flush()
	_, err := fmt.Fprintf(writer, "Id\tName\tPassword\tDescription\tWrite\n")
	if err != nil {
		return fmt.Errorf("failed to write to tabwriter: %w", err)
	}
	for _, user := range users {
		_, err = fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%t\n", user.Id, user.Username, user.Password, user.Description, user.Write)
		if err != nil {
			return fmt.Errorf("failed to write to tabwriter: %w", err)
		}
	}
	return nil
}
