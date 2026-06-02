package displaytext

import (
	"example.com/led-display-api/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "display-text",
}

func init() {
	Cmd.AddCommand(createDisplayTextCmd)
	root.AddCommand(Cmd)
}
