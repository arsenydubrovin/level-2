package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grep [-ABCFchinv] [pattern] [file]",
	Short: "Searches text patterns in file",
	Run:   runGrep,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntP("after", "A", 0, "Print N lines after each matching line")
	rootCmd.Flags().IntP("before", "B", 0, "Print N lines before each matching line")
	rootCmd.Flags().IntP("context", "C", 0, "Print N lines of output context")
	rootCmd.Flags().BoolP("count", "c", false, "Display only the count of matching lines")
	rootCmd.Flags().BoolP("fixed", "F", false, "Fixed string (exact match) rather than a pattern")
	rootCmd.Flags().BoolP("ignore-case", "i", false, "Ignore case when matching")
	rootCmd.Flags().BoolP("line-numbers", "n", false, "Print line numbers with output")
	rootCmd.Flags().BoolP("invert", "v", false, "Invert match (exclude lines that match the pattern)")
}
