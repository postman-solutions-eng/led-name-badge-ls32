package setupauth

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"os"
	"strconv"
	"strings"
)

func promptLine(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}
	return strings.TrimSpace(line), nil
}

func promptSecret(prompt string) (string, error) {
	fmt.Print(prompt)
	bytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		return "", fmt.Errorf("failed to read secret: %w", err)
	}
	return string(bytes), nil
}

var setupAuthInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup authentication credentials step by step",
	RunE: func(cmd *cobra.Command, args []string) error {
		type authOption struct {
			label string
			run   func() error
		}

		var options []authOption

		options = append(options, authOption{
			label: "API key",
			run: func() error {
				apiKey, err := promptSecret("Enter API key: ")
				if err != nil {
					return err
				}
				if apiKey == "" {
					return fmt.Errorf("API key is required")
				}
				return runApiKeyAuth(apiKey)
			},
		})

		if len(options) > 1 {
			fmt.Println("Select authentication type:")
			for i, opt := range options {
				fmt.Printf("  %d. %s\n", i+1, opt.label)
			}
			choiceStr, err := promptLine("Enter choice: ")
			if err != nil {
				return err
			}
			choice, err := strconv.Atoi(choiceStr)
			if err != nil || choice < 1 || choice > len(options) {
				return fmt.Errorf("invalid choice: %q", choiceStr)
			}
			return options[choice-1].run()
		}

		if len(options) == 0 {
			return fmt.Errorf("no authentication methods available")
		}
		return options[0].run()
	},
}

func init() {
	Cmd.AddCommand(setupAuthInitCmd)
}
