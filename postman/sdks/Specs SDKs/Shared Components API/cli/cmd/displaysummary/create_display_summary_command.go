package displaysummary

import (
	"encoding/json"
	"example.com/led-display-api/root"
	"example.com/led-display-api/sdk/displaysummary"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var createDisplaySummaryCmd = &cobra.Command{
	Use:   "create-display-summary",
	Short: "Display Summary",
	Long:  "Displays a summary message on the LED badge.\n\nWithout an API key (`X-API-Key` header or `apiKey` in the body), shows the built-in demo summary text. When a key is provided, the endpoint fetches services from the Postman API Catalog and renders a compact health summary on the badge.\n\nNote: The API does not validate the `type` parameter and will return success for any type value provided.\n\nBody schema:\n  {\n    \"type\": \"type\",\n    \"customText\": \"customText\",\n    \"apiKey\": \"apiKey\",\n    \"systemEnvironmentId\": \"e8f94f60-f018-425a-afdd-dfbec894def8\"\n  }\n\nExamples:\n  --body '{\"type\":\"type\",\"customText\":\"customText\",\"apiKey\":\"apiKey\",\"systemEnvironmentId\":\"e8f94f60-f018-425a-afdd-dfbec894def8\"}'\n  --body-file ./body.json",
	RunE: func(cmd *cobra.Command, args []string) error {
		bodyStr, err := cmd.Flags().GetString("body")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to get body param: %v\n", err)
			return err
		}
		bodyFile, err := cmd.Flags().GetString("body-file")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to get body-file param: %v\n", err)
			return err
		}

		var bodyContent []byte
		if bodyFile != "" {
			bodyContent, err = os.ReadFile(bodyFile)
			if err != nil {
				return err
			}
		} else {
			bodyContent = []byte(bodyStr)
		}

		var requestBody displaysummary.CreateDisplaySummaryRequest
		if len(bodyContent) > 0 {
			if err := json.Unmarshal(bodyContent, &requestBody); err != nil {
				return err
			}
		}

		client := root.CreateSdkClient()
		response, err := client.DisplaySummary.CreateDisplaySummary(cmd.Context(), requestBody)
		if err != nil {
			return err
		}

		jsonData, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			fmt.Println(response)
		} else {
			fmt.Println(string(jsonData))
		}

		return nil
	},
}

func init() {
	createDisplaySummaryCmd.Flags().String("body", "", "Request body as inline JSON, e.g. '{\"type\":\"type\",\"customText\":\"customText\",\"apiKey\":\"apiKey\",\"systemEnvironmentId\":\"e8f94f60-f018-425a-afdd-dfbec894def8\"}'")
	createDisplaySummaryCmd.Flags().String("body-file", "", "Path to a file containing the request body")
	createDisplaySummaryCmd.MarkFlagsMutuallyExclusive("body", "body-file")
}
