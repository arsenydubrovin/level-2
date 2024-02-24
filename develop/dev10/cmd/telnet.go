package cmd

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/spf13/cobra"
)

func runTelnet(cmd *cobra.Command, args []string) {
	var address string

	if len(args) == 2 {
		address = fmt.Sprintf("%s:%s", args[0], args[1])
	} else if len(args) == 1 {
		address = fmt.Sprintf("%s:%d", args[0], 23)
	} else {
		writeError(cmd.Usage())
	}

	t, err := cmd.Flags().GetDuration("timeout")
	if err != nil {
		writeError(err)
	}

	conn, err := net.DialTimeout("tcp", address, t)
	if err != nil {
		writeError(err)
	}
	defer conn.Close()

	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			writeError(err)
		}
	}()

	_, err = io.Copy(conn, os.Stdin)
	if err != nil {
		writeError(err)
	}
}

// writeError выводит err в stderr.
func writeError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
