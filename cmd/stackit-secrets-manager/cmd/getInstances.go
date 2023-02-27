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

var getInstancesCmd = &cobra.Command{
	Use:   "instances",
	Short: "Returns all instances from the Secrets Manager.",
	Long:  `Returns all instances from the Secrets Manager.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := getInstances(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	},
}

func init() {
	getCmd.AddCommand(getInstancesCmd)
}

func getInstances() error {
	client, err := createClient()
	if err != nil {
		return err
	}
	response, err := client.GetV1ProjectsProjectIdInstances(context.Background(), projectId)
	if err != nil {
		return fmt.Errorf("failed to get instances: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}

	var instances api.InstanceList
	if err := json.NewDecoder(response.Body).Decode(&instances); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}
	if err := printInstances(instances.Instances); err != nil {
		return err
	}
	return nil
}

func printInstances(instances []api.Instance) error {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer writer.Flush()
	_, err := fmt.Fprintf(writer, "Id\tName\tUsers\tSecrets\tAPI URL\tSecrets Engine\tState\n")
	if err != nil {
		return fmt.Errorf("failed to write to tabwriter: %w", err)
	}
	for _, instance := range instances {
		_, err = fmt.Fprintf(writer, "%s\t%s\t%d\t%d\t%s\t%s\t%s\n", instance.Id, instance.Name, instance.UserLimit, instance.SecretLimit, instance.ApiUrl, instance.SecretsEngine, instance.State)
		if err != nil {
			return fmt.Errorf("failed to write to tabwriter: %w", err)
		}
	}
	return nil
}
