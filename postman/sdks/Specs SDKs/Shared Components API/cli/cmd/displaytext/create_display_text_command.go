package displaytext

import (
	"encoding/json"
	"example.com/led-display-api/root"
	"example.com/led-display-api/sdk/displaytext"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var createDisplayTextCmd = &cobra.Command{
	Use:   "create-display-text",
	Short: "Display Text",
	Long:  "Updates the text and visual content displayed on the LED name badge. Accepts text and icon codes in the format `:icon_name:`.\n\n<img src=\"https://content.pstmn.io/05b1ef2c-9fd3-4f0e-8b2f-76998fb3f6e5/aW1hZ2UucG5n\" width=\"240\">\n\n## Supported Icons\n\nThe following icon codes can be embedded in text using the `:icon_name:` syntax:\n\n- `:ball:` - Filled circle\n    \n- `:happy:` - Simple smiley face\n    \n- `:happy2:` - Larger smiley (2 columns wide)\n    \n- `:heart:` - Outline heart\n    \n- `:HEART:` - Filled heart\n    \n- `:heart2:` - Larger outline heart (2 columns wide)\n    \n- `:HEART2:` - Larger filled heart (2 columns wide)\n    \n- `:fablab:` - FabLab logo\n    \n- `:bicycle:` - Bicycle icon (3 columns wide)\n    \n- `:bicycle_r:` - Bicycle facing right (3 columns wide)\n    \n- `:owncloud:` - OwnCloud logo (3 columns wide)\n    \n- `:octocat:` - GitHub Octocat\n    \n- `:smile:` - Smile emoji\n    \n- `:star:` - Star icon\n    \n- `:sun:` - Sun icon\n    \n\n## Character Restrictions\n\n**Supported characters:**\n\n- Letters: A-Z, a-z\n    \n- Numbers: 0-9\n    \n- Special characters: `^ !\"$%&/()=?` °}\\]\\[{@ \\~ |<>,;.:-_#'+_\\`\n    \n- German umlauts: äöüÄÖÜß\n    \n- French/European accents: àäòöùüèéêëôöûîïÿçÀÅÄÉÈÊËÖÔÜÛÙŸŠ\n    \n\n**NOT supported:**\n\n- Unicode emoji (e.g., 🌍) - will return a 400 error\n\nBody schema:\n  {\n    \"text\": \"text\"\n  }\n\nExamples:\n  --body '{\"text\":\"text\"}'\n  --body-file ./body.json",
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

		var requestBody displaytext.CreateDisplayTextRequest
		if len(bodyContent) > 0 {
			if err := json.Unmarshal(bodyContent, &requestBody); err != nil {
				return err
			}
		}

		client := root.CreateSdkClient()
		response, err := client.DisplayText.CreateDisplayText(cmd.Context(), requestBody)
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
	createDisplayTextCmd.Flags().String("body", "", "Request body as inline JSON, e.g. '{\"text\":\"text\"}'")
	createDisplayTextCmd.Flags().String("body-file", "", "Path to a file containing the request body")
	createDisplayTextCmd.MarkFlagsMutuallyExclusive("body", "body-file")
}
