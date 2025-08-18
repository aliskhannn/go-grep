package main

import (
	"fmt"
	"github.com/aliskhannn/go-grep/internal/grep"
	"github.com/spf13/pflag"
	"os"
)

func main() {
	// Initialize command-line flags for the grep utility.
	flags := grep.InitFlags()
	pflag.Parse()

	args := pflag.Args()
	if len(args) == 0 {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: gogrep [OPTIONS] PATTERN [FILE...]")
		os.Exit(1)
	}

	pattern := args[0]
	files := args[1:]

	// If no files are specified, default to reading from stdin.
	if len(files) == 0 {
		files = []string{"-"}
	}

	// Create a configuration for the grep operation using the parsed flags and arguments.
	cfg := grep.Config{
		Pattern:     pattern,
		Files:       files,
		After:       *flags.After,
		Before:      *flags.Before,
		Context:     *flags.Context,
		CountOnly:   *flags.Count,
		IgnoreCase:  *flags.Ignore,
		InvertMatch: *flags.Invert,
		Fixed:       *flags.Fixed,
		LineNum:     *flags.LineNum,
	}

	// Run the grep operation with the provided configuration and output the results.
	if err := grep.Run(cfg, os.Stdout, os.Stderr); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(2)
	}
}
