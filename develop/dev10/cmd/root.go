package cmd

import (
	"os"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "telnet [host] [port] [flags]",
	Short: "Establishes a connection with a remote host",
	Run:   runTelnet,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().DurationP("timeout", "t", 10*time.Second, "Connection timeout")
}
