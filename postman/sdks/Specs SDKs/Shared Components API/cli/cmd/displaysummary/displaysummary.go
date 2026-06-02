package displaysummary

import (
	"example.com/led-display-api/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "display-summary",
}

func init() {
	Cmd.AddCommand(createDisplaySummaryCmd)
	root.AddCommand(Cmd)
}
