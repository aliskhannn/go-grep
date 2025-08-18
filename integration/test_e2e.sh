#!/bin/bash
set -euo pipefail

go build -o gogrep ./cmd/gogrep

echo "== Test 1: simple match"
./gogrep foo testdata/sample.txt > out1.txt
grep foo testdata/sample.txt > ref1.txt
diff -u ref1.txt out1.txt

echo "== Test 2: ignore case"
./gogrep -i foo testdata/sample.txt > out2.txt
grep -i foo testdata/sample.txt > ref2.txt
diff -u ref2.txt out2.txt

echo "All tests passed!"