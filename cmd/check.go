package cmd

import (
	"vimes/internal"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Run a check against the known recorded baseline",
	Run: func(cmd *cobra.Command, args []string) {
		for _, path := range config.Paths {
			println(path)
		}

		ff := internal.NewFileFinder()
		ff.FindFiles()
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
