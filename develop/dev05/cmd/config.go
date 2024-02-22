package cmd

import "github.com/spf13/cobra"

type config struct {
	after       int
	before      int
	context     int
	count       bool
	fixed       bool
	ignoreCase  bool
	lineNumbers bool
	invert      bool
}

// readConfig читает конфигурацию из флагов
func readConfig(cmd *cobra.Command) (*config, error) {
	after, err := cmd.Flags().GetInt("after")
	if err != nil {
		return nil, err
	}

	before, err := cmd.Flags().GetInt("before")
	if err != nil {
		return nil, err
	}

	context, err := cmd.Flags().GetInt("context")
	if err != nil {
		return nil, err
	}

	count, err := cmd.Flags().GetBool("count")
	if err != nil {
		return nil, err
	}

	fixed, err := cmd.Flags().GetBool("fixed")
	if err != nil {
		return nil, err
	}

	ignoreCase, err := cmd.Flags().GetBool("ignore-case")
	if err != nil {
		return nil, err
	}

	lineNumbers, err := cmd.Flags().GetBool("line-numbers")
	if err != nil {
		return nil, err
	}

	invert, err := cmd.Flags().GetBool("invert")
	if err != nil {
		return nil, err
	}

	return &config{
		after:       after,
		before:      before,
		context:     context,
		count:       count,
		fixed:       fixed,
		ignoreCase:  ignoreCase,
		lineNumbers: lineNumbers,
		invert:      invert,
	}, nil
}
