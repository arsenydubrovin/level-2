package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func commandLs(args []string) error {
	if len(args) != 2 {
		return errors.New("usage: ls [dir]")
	}

	cmd := exec.Command("ls", args[1])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func commandCd(args []string) error {
	if len(args) != 2 {
		return errors.New("usage: cd [dir]")
	}

	err := os.Chdir(args[1])
	if err != nil {
		return err
	}

	return nil
}

func commandPwd(args []string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, dir)

	return nil
}

func commandEcho(args []string) error {
	if len(args) < 2 {
		return errors.New("usage: echo [args...]")
	}

	fmt.Fprintln(os.Stdout, strings.Join(args[1:], " "))

	return nil
}

func commandKill(args []string) error {
	if len(args) != 2 {
		return errors.New("usage: kill [pid]")
	}

	cmd := exec.Command("kill", args[1])

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func commandPs(args []string) error {
	if len(args) != 1 {
		return errors.New("usage: ps")
	}

	cmd := exec.Command("ps", "aux")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func commandExit(args []string) error {
	fmt.Fprintln(os.Stdout, "exiting...")
	os.Exit(0)

	return nil
}
