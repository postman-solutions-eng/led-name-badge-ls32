package setupauth

import (
	"example.com/led-display-api/root"
	"github.com/spf13/cobra"
)

func SetupAuth(client any) {
	if setter, ok := client.(ApiKeyAuthSetter); ok {
		ConfigureApiKeyAuth(setter)
	}
}

var Cmd = &cobra.Command{
	Use:   "setup-auth",
	Short: "Configure authentication credentials",
}

func init() {
	root.AddCommand(Cmd)
	root.RegisterAuthConfigurator(SetupAuth)
}
