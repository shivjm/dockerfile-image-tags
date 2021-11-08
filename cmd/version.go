package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var Version = "development"
var Revision = ""

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Prints version",
		Long:  "Prints the program’s version information.",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			git := ""
			r := strings.TrimSpace(Revision)

			if len(r) > 0 {
				git = fmt.Sprintf(" (r%s)", r)
			}

			fmt.Printf("%s %s%s", ProgramName, Version, git)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
