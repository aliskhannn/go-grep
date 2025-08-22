# GoGrep

GoGrep is a Go implementation of a simplified UNIX `grep` utility. It allows filtering lines from text files or standard
input based on a pattern, with support for various flags including context, line numbers, case-insensitive search,
fixed-string matching, and counting matches.

---

## Features

* Search for a pattern in files or stdin.
* Support for regular expressions or fixed string matching (`-F`).
* Case-insensitive search (`-i`).
* Invert match (`-v`).
* Show line numbers (`-n`).
* Show context lines before (`-B`) or after (`-A`) matches, or both (`-C`).
* Count only matches (`-c`).
* Combination of multiple flags is supported.

---

## Project Structure

```
gogrep/
├── cmd/
│   └── gogrep/
│       └── main.go          # Entry point of the application
├── internal/
│   └── grep/
│       ├── flags.go         # Flag parsing and validation
│       ├── grep.go          # Core grep functionality
│       └── grep_test.go     # Unit tests
├── integration/
│   └── test_e2e.sh          # End-to-end tests
├── testdata/
│   └── sample.txt           # Sample input file for tests
├── go.mod
├── go.sum
├── Makefile                 # Build and test commands
└── README.md
```

---

## Installation

```bash
# Clone the repository
git clone https://github.com/aliskhannn/go-grep.git
cd go-grep

# Build the binary
make build
```

---

## Usage

```bash
# Search for "foo" in a file
./bin/gogrep foo testdata/sample.txt

# Case-insensitive search
./bin/gogrep -i foo testdata/sample.txt

# Show line numbers
./bin/gogrep -n foo testdata/sample.txt

# Show 2 lines after each match
./bin/gogrep -A 2 foo testdata/sample.txt

# Show 1 line before and after each match
./bin/gogrep -C 1 foo testdata/sample.txt

# Invert match (show lines that do NOT contain the pattern)
./bin/gogrep -v foo testdata/sample.txt

# Count the number of matching lines
./bin/gogrep -c -i foo testdata/sample.txt

# Search using a fixed string instead of regex
./bin/gogrep -F "error: something bad" testdata/sample.txt
```

---

## Testing

### Unit tests

```bash
make test
```

### Integration / end-to-end tests

```bash
make integration
```

---