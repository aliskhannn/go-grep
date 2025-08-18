package grep

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// Config describes the utility's operating parameters.
type Config struct {
	Pattern     string   // the pattern to search for
	Files       []string // list of files to search in (or "-" for stdin)
	After       int      // number of lines to print after a match
	Before      int      // number of lines to print before a match
	Context     int      // number of lines to print before and after a match
	CountOnly   bool     // print only the count of matching lines
	IgnoreCase  bool     // ignore case distinctions in patterns and data
	InvertMatch bool     // print lines that do not match the pattern
	Fixed       bool     // use fixed strings for matching (no regex)
	LineNum     bool     // print line numbers with output lines
}

// Run executes the grep functionality based on the provided configuration.
func Run(cfg Config, stdout io.Writer, stderr io.Writer) error {
	matcher, err := buildMatcher(cfg)
	if err != nil {
		return fmt.Errorf("invalid pattern: %w", err)
	}

	totalMatches := 0

	for _, fname := range cfg.Files {
		var r io.Reader
		var file *os.File

		if fname == "-" {
			r = os.Stdin
		} else {
			file, err = os.Open(fname)
			if err != nil {
				return fmt.Errorf("failed to open file %s: %w", fname, err)
			}

			r = file
		}

		matches, err := processFile(r, stdout, cfg, matcher)
		if err != nil {
			return err
		}

		// Close the file if it was opened.
		if file != nil {
			_ = file.Close()
		}

		totalMatches += matches
	}

	if cfg.CountOnly {
		_, _ = fmt.Fprintln(stdout, totalMatches)
	}

	return nil
}

// buildMatcher creates a function that checks whether a string matches a template.
func buildMatcher(cfg Config) (func(string) bool, error) {
	pattern := cfg.Pattern
	if cfg.IgnoreCase {
		pattern = "(?i)" + pattern // make the pattern case-insensitive
	}

	if cfg.Fixed {
		// If fixed strings are used, we can use a simple contains check.
		if cfg.IgnoreCase {
			return func(s string) bool {
				return strings.Contains(strings.ToLower(s), strings.ToLower(cfg.Pattern))
			}, nil
		}

		return func(s string) bool {
			return strings.Contains(s, cfg.Pattern)
		}, nil
	}

	// Compile the regular expression pattern.
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	return re.MatchString, nil
}

// processFile reads a file, applies the matcher to each line,
// and writes the results to the output.
func processFile(r io.Reader, w io.Writer, cfg Config, matcher func(string) bool) (int, error) {
	scanner := bufio.NewScanner(r)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	matchCount := 0
	toPrint := make(map[int]bool)

	// Iterate through the lines and apply the matcher.
	for i, line := range lines {
		matched := matcher(line)

		if cfg.InvertMatch {
			matched = !matched
		}

		if matched {
			matchCount++

			if cfg.CountOnly {
				continue
			}

			start := max(0, i-max(cfg.Before, cfg.Context))
			end := min(len(lines)-1, i+max(cfg.After, cfg.Context))
			for j := start; j <= end; j++ {
				toPrint[j] = true
			}
		}
	}

	// Print the count of matches if requested.
	if !cfg.CountOnly {
		for i := 0; i < len(lines); i++ {
			if toPrint[i] {
				if cfg.LineNum {
					_, _ = fmt.Fprintf(w, "%d:%s\n", i+1, lines[i])
				} else {
					_, _ = fmt.Fprintln(w, lines[i])
				}
			}
		}
	}

	return matchCount, nil
}
