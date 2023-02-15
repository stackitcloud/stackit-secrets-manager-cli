package cmd

import (
	"context"
	"fmt"
	"github.com/stackitcloud/stackit-secrets-manager-cli/internal/api"
	"net/http"
)

func createClient() (*api.Client, error) {
	apiClient, err := api.NewClient(secretsManagerApiUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create api client: %v", err)
	}
	apiClient.RequestEditors = []api.RequestEditorFn{
		func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", authenticationToken))
			return nil
		},
	}
	return apiClient, nil
}
