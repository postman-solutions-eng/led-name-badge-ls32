package predefinedicons

import (
	"encoding/json"
	"example.com/led-display-api/root"
	"fmt"
	"github.com/spf13/cobra"
)

var getPredefinedIconsCmd = &cobra.Command{
	Use:   "get-predefined-icons",
	Short: "Get Predefined Icons",
	RunE: func(cmd *cobra.Command, args []string) error {

		client := root.CreateSdkClient()
		response, err := client.PredefinedIcons.GetPredefinedIcons(cmd.Context())
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
}
