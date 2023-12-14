package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-secrets-manager-cli/internal/api"
)

var (
	updateAclInstanceId string
	updateAclAclId      string
	updateAclCidr       string
)

var updateAclCmd = &cobra.Command{
	Use:   "acl",
	Short: "Updates an acl for the Secrets Manager.",
	Long:  `Updates an acl for the Secrets Manager.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := updateAcl(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	},
}

func init() {
	updateCmd.AddCommand(updateAclCmd)

	updateAclCmd.PersistentFlags().StringVar(&updateAclInstanceId, "instance-id", "", "The UUID of the instance to update the acl for.")
	_ = updateAclCmd.MarkPersistentFlagRequired("instance-id")
	updateAclCmd.PersistentFlags().StringVar(&updateAclAclId, "acl-id", "", "The UUID of the acl to update the cidr from.")
	_ = updateAclCmd.MarkPersistentFlagRequired("acl-id")
	updateAclCmd.PersistentFlags().StringVar(&updateAclCidr, "cidr", "", "The updated cidr.")
}

func updateAcl() error {
	client, err := createClient()
	if err != nil {
		return err
	}
	request := api.PutV1ProjectsProjectIdInstancesInstanceIdAclsAclIdJSONRequestBody{
		Cidr: updateAclCidr,
	}
	response, err := client.PutV1ProjectsProjectIdInstancesInstanceIdAclsAclId(context.Background(), projectId, updateAclInstanceId, updateAclAclId, request)
	if err != nil {
		return fmt.Errorf("failed to update acl: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}
	fmt.Printf("Acl %s updated\n", updateAclAclId)
	return nil
}
