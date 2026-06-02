package setupauth

import (
	"example.com/led-display-api/internal/credentials"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"os"
)

type ApiKeyAuthSetter interface {
	SetAPIKey(apiKey string)
}

func ConfigureApiKeyAuth(client ApiKeyAuthSetter) {
	creds, err := credentials.LoadCredentials()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to load credentials: %v\n", err)
		return
	}
	if creds != nil && creds.ApiKey != "" {
		client.SetAPIKey(creds.ApiKey)
	}
}

func runApiKeyAuth(apiKey string) error {
	creds, err := credentials.LoadCredentials()
	if err != nil {
		return fmt.Errorf("failed to load existing credentials: %w", err)
	}
	if creds == nil {
		creds = &credentials.Credentials{}
	}
	creds.ApiKey = apiKey
	if err := credentials.SaveCredentials(*creds); err != nil {
		return fmt.Errorf("failed to save credentials: %w", err)
	}
	fmt.Println("API key authentication credentials saved successfully.")
	return nil
}

var setupAuthApiKeyCmd = &cobra.Command{
	Use:   "api-key [value]",
	Short: "Configure API key authentication",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var apiKey string
		if len(args) > 0 {
			apiKey = args[0]
		} else {
			fmt.Print("API Key: ")
			byteKey, err := term.ReadPassword(int(os.Stdin.Fd()))
			fmt.Println()
			if err != nil {
				return fmt.Errorf("failed to read api key: %w", err)
			}
			apiKey = string(byteKey)
		}

		return runApiKeyAuth(apiKey)
	},
}

func init() {
	Cmd.AddCommand(setupAuthApiKeyCmd)
}
