package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/stackitcloud/stackit-secrets-manager-cli/internal/api"

	"github.com/spf13/cobra"
)

var (
	getAclsInstanceId string
)

var getAclsCmd = &cobra.Command{
	Use:   "acls",
	Short: "Returns all acls from the Secrets Manager.",
	Long:  `Returns all acls from the Secrets Manager.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := getAcls(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	},
}

func init() {
	getCmd.AddCommand(getAclsCmd)

	getAclsCmd.PersistentFlags().StringVar(&getAclsInstanceId, "instance-id", "", "The UUID of the instance to get the acls for.")
	_ = getAclsCmd.MarkPersistentFlagRequired("instance-id")
}

func getAcls() error {
	client, err := createClient()
	if err != nil {
		return err
	}
	response, err := client.GetV1ProjectsProjectIdInstancesInstanceIdAcls(context.Background(), projectId, getAclsInstanceId)
	if err != nil {
		return fmt.Errorf("failed to get acls: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}

	var acls api.AclList
	if err := json.NewDecoder(response.Body).Decode(&acls); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}
	if err := printAcls(acls.Acls); err != nil {
		return err
	}
	return nil
}

func printAcls(acls []api.Acl) error {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer writer.Flush()
	_, err := fmt.Fprintf(writer, "Id\tCidr\n")
	if err != nil {
		return fmt.Errorf("failed to write to tabwriter: %w", err)
	}
	for _, acl := range acls {
		_, err = fmt.Fprintf(writer, "%s\t%s\n", acl.Id, acl.Cidr)
		if err != nil {
			return fmt.Errorf("failed to write to tabwriter: %w", err)
		}
	}
	return nil
}
