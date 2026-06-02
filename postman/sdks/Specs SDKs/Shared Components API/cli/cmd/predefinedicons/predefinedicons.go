package predefinedicons

import (
	"example.com/led-display-api/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "predefined-icons",
}

func init() {
	Cmd.AddCommand(getPredefinedIconsCmd)
	root.AddCommand(Cmd)
}
