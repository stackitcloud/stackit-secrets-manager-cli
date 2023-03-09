package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-secrets-manager-cli/internal/api"
	"net/http"
)

var (
	createInstanceName string
)

var createInstanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "Creates a new instance for the Secrets Manager.",
	Long:  `Creates a new instance for the Secrets Manager.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := createInstance(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	},
}

func init() {
	createCmd.AddCommand(createInstanceCmd)

	createInstanceCmd.PersistentFlags().StringVar(&createInstanceName, "name", "", "The name to set for the instance.")
	_ = createInstanceCmd.MarkPersistentFlagRequired("name")
}

func createInstance() error {
	client, err := createClient()
	if err != nil {
		return err
	}
	request := api.PostV1ProjectsProjectIdInstancesJSONRequestBody{
		Name: createInstanceName,
	}
	response, err := client.PostV1ProjectsProjectIdInstances(context.Background(), projectId, request)
	if err != nil {
		return fmt.Errorf("failed to create instance: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}

	var instance api.Instance
	if err := json.NewDecoder(response.Body).Decode(&instance); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}
	if err := printInstances([]api.Instance{instance}); err != nil {
		return err
	}
	return nil
}
