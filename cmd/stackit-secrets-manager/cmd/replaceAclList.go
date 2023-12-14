package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-secrets-manager-cli/internal/api"
)

var (
	updateAclsInstanceId string
	updateAclsCidrString []string
)

var updateAclsCmd = &cobra.Command{
	Use:   "acl",
	Short: "Replaces an acl list with another acl list.",
	Long:  `Replaces an acl list with another acl list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := updateAcl(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	},
}

func init() {
	updateCmd.AddCommand(updateAclsCmd)
	updateAclsCmd.PersistentFlags().StringVar(&updateAclsInstanceId, "instance-id", "", "The UUID of the instance to update the acls.")
}

func replaceAclList() error {
	client, err := createClient()
	if err != nil {
		return err
	}
	var updateAclsCidr []api.AclUpdate
	for _, cidr := range updateAclsCidrString {
		updateAclsCidr = append(updateAclsCidr, api.AclUpdate{
			Cidr: cidr,
		})
	}

	request := api.PutV1ProjectsProjectIdInstancesInstanceIdAclsJSONRequestBody{
		Cidrs: &updateAclsCidr,
	}
	response, err := client.PutV1ProjectsProjectIdInstancesInstanceIdAcls(context.Background(), projectId, updateAclsInstanceId, request)
	if err != nil {
		return fmt.Errorf("failed to update acls %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}
	fmt.Printf("Acls were updated")
	return nil
}