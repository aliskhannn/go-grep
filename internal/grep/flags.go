package grep

import "github.com/spf13/pflag"

// Flags stores pointers to the flag values returned by pflag.
type Flags struct {
	After   *int
	Before  *int
	Context *int
	Count   *bool
	Ignore  *bool
	Invert  *bool
	Fixed   *bool
	LineNum *bool
}

// InitFlags initializes and returns a Flags struct with command-line flags for grep-like functionality.
func InitFlags() Flags {
	return Flags{
		After:   pflag.IntP("after", "A", 0, "Print N lines after a match (context)"),
		Before:  pflag.IntP("before", "B", 0, "Print N lines before a match (context)"),
		Context: pflag.IntP("around", "C", 0, "Print N lines before and after a match (equivalent to -A N -B N)"),
		Count:   pflag.BoolP("count", "c", false, "Print only a count of matching lines"),
		Ignore:  pflag.BoolP("ignore-case", "i", false, "Ignore case distinctions in patterns and data"),
		Invert:  pflag.BoolP("invert-match", "v", false, "Print lines that do not match the pattern"),
		Fixed:   pflag.BoolP("fixed-strings", "F", false, "Use fixed strings for matching instead of regular expressions"),
		LineNum: pflag.BoolP("line-number", "n", false, "Print line numbers with output lines"),
	}
}
