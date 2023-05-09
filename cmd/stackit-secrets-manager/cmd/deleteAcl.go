package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	deleteAclInstanceId string
)

var deleteAclCmd = &cobra.Command{
	Use:   "acls <aclid>",
	Short: "Deletes an acl from the Secrets Manager.",
	Long:  `Deletes an acl from the Secrets Manager.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, aclId := range args {
			if err := deleteAcl(aclId); err != nil {
				fmt.Printf("ERROR: %v\n", err)
			}
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteAclCmd)

	deleteAclCmd.PersistentFlags().StringVar(&deleteAclInstanceId, "instance-id", "", "The UUID of the instance to delete the acls from.")
	_ = deleteAclCmd.MarkPersistentFlagRequired("instance-id")
}

func deleteAcl(aclId string) error {
	client, err := createClient()
	if err != nil {
		return err
	}
	response, err := client.DeleteV1ProjectsProjectIdInstancesInstanceIdAclsAclId(context.Background(), projectId, deleteAclInstanceId, aclId)
	if err != nil {
		return fmt.Errorf("failed to delete acl: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}
	fmt.Printf("Acl %s deleted\n", aclId)
	return nil
}
