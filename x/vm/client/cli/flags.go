package cli

import "github.com/spf13/cobra"

const (
	FlagOutput = "to-file"
)

func AddOutputFlagToCmd(cmd *cobra.Command) {
	cmd.Flags().String(FlagOutput, "script.json", "Output file path")
}
