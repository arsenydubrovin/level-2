package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sort",
	Short: "Sorts the lines in the specified file",
	Long: `Sorts the lines in the specified file.

It can sort by the specified column, by numeric value
or in reverse order. You can disable printing duplicate lines`,
	Run: runSort,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	sortColumn   int
	numericSort  bool
	reverseOrder bool
	uniqueLines  bool
)

func init() {
	rootCmd.Flags().IntVarP(&sortColumn, "column", "k", 1, "Column number for sorting")
	rootCmd.Flags().BoolVarP(&numericSort, "numeric", "n", false, "Enable numeric sorting")
	rootCmd.Flags().BoolVarP(&reverseOrder, "reverse", "r", false, "Sort in reverse order")
	rootCmd.Flags().BoolVarP(&uniqueLines, "unique", "u", false, "Display only unique lines")
}
