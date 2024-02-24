package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var cmds = map[string]func([]string) error{
	"ls":   commandLs,
	"cd":   commandCd,
	"pwd":  commandPwd,
	"echo": commandEcho,
	"kill": commandKill,
	"ps":   commandPs,
	"exit": commandExit,
}

func runShell() {
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "→ ")

		input, err := r.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Fprintln(os.Stdout, "")
				os.Exit(0)
			}

			fmt.Fprintf(os.Stderr, "can't read the input: %s\n", err)
			continue
		}

		input = strings.TrimSuffix(input, "\n")
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		cmd, ok := cmds[args[0]]
		if !ok {
			fmt.Fprintf(os.Stderr, "unknown command: %s\n", args[0])
			continue
		}

		err = cmd(args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
