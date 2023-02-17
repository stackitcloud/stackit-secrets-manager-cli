package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/term"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Creates the configuration file needed to interact with the Secrets Manager API.",
	Long:  `Creates the configuration file needed to interact with the Secrets Manager API.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		// For the configure command, we cannot have the authentication token and project id set as required flags.
		// As that would prevent the configure command from running when no config file was written yet. As cobra does
		// not provide an inverse for MarkPersistentFlagRequired(), we set the required annotation on the root command
		// ourselves.
		_ = rootCmd.PersistentFlags().SetAnnotation("authentication-token", cobra.BashCompOneRequiredFlag, []string{"false"})
		_ = rootCmd.PersistentFlags().SetAnnotation("project-id", cobra.BashCompOneRequiredFlag, []string{"false"})
	},
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReaderSize(os.Stdin, 32*1024)

		fmt.Printf("Authentication Token [%s]: ", authenticationToken)
		input, err := readLongString()
		if err != nil {
			fmt.Printf("ERROR: Failed to read user input: %v\n", err)
			return
		}
		input = strings.TrimSpace(input)
		if input != "" {
			viper.Set("authentication-token", input)
		}

		fmt.Printf("Project UUID [%s]: ", projectId)
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf("ERROR: Failed to read user input: %v\n", err)
			return
		}
		input = strings.TrimSpace(input)
		if input != "" {
			viper.Set("project-id", input)
		}

		// Viper has a known issue with WriteConfig running into an error in case the config file does not exist. Therefore,
		// we first try SafeWriteConfig which only works in cases where the config file does not exist. If that
		// fails, the config file is probably already there, and we use WriteConfig.
		if err := viper.SafeWriteConfig(); err != nil {
			if err := viper.WriteConfig(); err != nil {
				fmt.Printf("ERROR: Failed to write config: %v\n", err)
				return
			}
		}
		fmt.Println("Configuration successfully written")
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}

// readLongString provides a way to read more than 1024 characters from the terminal by switching the terminal into
// raw mode. Otherwise, long strings like the authentication token would be truncated to 1024 characters because of
// canonical input mode for terminals.
func readLongString() (string, error) {
	termStateBackup, _ := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), termStateBackup) // nolint:errcheck

	result := ""
	characters := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(characters)
		if err != nil {
			return result, err
		}
		for i := 0; i < n; i++ {
			if characters[i] == 0x03 || // Ctrl+C
				characters[i] == 0x0d { // Enter
				fmt.Print("\r\n")
				return result, nil
			}
			fmt.Printf("%s", string(characters[i]))
			result = result + string(characters[i])
		}
	}
}
