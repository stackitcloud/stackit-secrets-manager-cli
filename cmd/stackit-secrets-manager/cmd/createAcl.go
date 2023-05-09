package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stackitcloud/stackit-secrets-manager-cli/internal/api"

	"github.com/spf13/cobra"
)

var (
	createAclInstanceId string
	createAclCidr       string
)

var createAclCmd = &cobra.Command{
	Use:   "acl",
	Short: "Creates a new acl for the Secrets Manager.",
	Long:  `Creates a new acl for the Secrets Manager.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := createAcl(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	},
}

func init() {
	createCmd.AddCommand(createAclCmd)

	createAclCmd.PersistentFlags().StringVar(&createAclInstanceId, "instance-id", "", "The UUID of the instance to create the acl for.")
	_ = createAclCmd.MarkPersistentFlagRequired("instance-id")
	createAclCmd.PersistentFlags().StringVar(&createAclCidr, "cidr", "", "The cidr for the acl.")
}

func createAcl() error {
	client, err := createClient()
	if err != nil {
		return err
	}
	request := api.PostV1ProjectsProjectIdInstancesInstanceIdAclsJSONRequestBody{
		Cidr: createAclCidr,
	}
	response, err := client.PostV1ProjectsProjectIdInstancesInstanceIdAcls(context.Background(), projectId, createAclInstanceId, request)
	if err != nil {
		return fmt.Errorf("failed to create acl: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}

	var acl api.Acl
	if err := json.NewDecoder(response.Body).Decode(&acl); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}
	if err := printAcls([]api.Acl{acl}); err != nil {
		return err
	}
	return nil
}
