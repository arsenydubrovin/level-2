package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cut [flags]",
	Short: "Extract selected fields from stdin",
	Run:   runCut,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	fields    string
	delimiter string
	separated bool
)

func init() {
	rootCmd.Flags().StringVarP(&fields, "fields", "f", "", "choose specific columns")
	rootCmd.MarkFlagRequired("fields")
	rootCmd.Flags().StringVarP(&delimiter, "delimiter", "d", "\t", "set the field delimiter character")
	rootCmd.Flags().BoolVarP(&separated, "separated", "s", false, "suppress lines with no delimiter character")
}
