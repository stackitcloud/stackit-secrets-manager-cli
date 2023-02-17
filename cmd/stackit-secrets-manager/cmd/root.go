package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

var (
	secretsManagerApiUrl string
	authenticationToken  string
	projectId            string
)

var rootCmd = &cobra.Command{
	Use:   "secrets-manager",
	Short: "A command line interface for interacting with the Secrets Manager API.",
	Long:  `A command line interface for interacting with the Secrets Manager API.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&secretsManagerApiUrl, "secrets-manager-api-url", "https://secrets-manager.api.eu01.stackit.cloud", "The url to the Secrets Manager API.")
	rootCmd.PersistentFlags().StringVar(&authenticationToken, "authentication-token", "", "The JWT token for authenticating with the Secrets Manager API.")
	_ = rootCmd.MarkPersistentFlagRequired("authentication-token")
	rootCmd.PersistentFlags().StringVar(&projectId, "project-id", "", "The project UUID the Secrets Manager resources are contained.")
	_ = rootCmd.MarkPersistentFlagRequired("project-id")
}

func initConfig() {
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
	viper.SetConfigName(".stackit-secrets-manager")
	viper.SetConfigType("yaml")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	viper.AddConfigPath(homeDir)

	// silently ignore missing config file, the user might provide all values on the command line or via
	// environment variables
	_ = viper.ReadInConfig()

	// There is some issue, where the integration of Cobra with Viper will result in wrong values, therefore we are
	// setting the values from viper manually. The issue is, that with the standard integration, viper will see, that
	// Cobra parameters are set - even if the command line parameter was not used and the default value was set. But
	// when Viper notices that the value is set, it will not overwrite the default value with the environment variable.
	// Another possibility would be to not have any default values set for cobra command line parameters, but this would
	// break the automatic help output from the cli. The manual way here seems the best solution for now.
	rootCmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		if !f.Changed && viper.IsSet(f.Name) {
			if err := rootCmd.PersistentFlags().Set(f.Name, fmt.Sprint(viper.Get(f.Name))); err != nil {
				log.Fatalf("unable to set value for command line parameter: %v", err)
			}
		}
	})
}
