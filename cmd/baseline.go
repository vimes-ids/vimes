package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var baselineCmd = &cobra.Command{
	Use:   "baseline",
	Short: "Record the current state of the system as the baseline for future checks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("baseline called")
	},
}

func init() {
	rootCmd.AddCommand(baselineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// baselineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// baselineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
