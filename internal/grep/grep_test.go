package grep

import (
	"bytes"
	"strings"
	"testing"
)

func TestFixedMatch(t *testing.T) {
	cfg := Config{
		Pattern: "foo",
		Fixed:   true,
	}
	matcher, _ := buildMatcher(cfg)
	if !matcher("foo bar") {
		t.Error("expected match for 'foo bar'")
	}
	if matcher("bar baz") {
		t.Error("did not expect match for 'bar baz'")
	}
}

func TestCaseInsensitive(t *testing.T) {
	cfg := Config{
		Pattern:    "foo",
		Fixed:      true,
		IgnoreCase: true,
	}
	matcher, _ := buildMatcher(cfg)
	if !matcher("FOO bar") {
		t.Error("expected case-insensitive match")
	}
}

func TestProcessFileSimple(t *testing.T) {
	input := "hello\nfoo\nworld\nfoo bar\n"
	r := strings.NewReader(input)
	var buf bytes.Buffer
	cfg := Config{
		Pattern: "foo",
		Fixed:   true,
	}
	matcher, _ := buildMatcher(cfg)
	count, err := processFile(r, &buf, cfg, matcher)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Errorf("expected 2 matches, got %d", count)
	}
	out := buf.String()
	if !strings.Contains(out, "foo") || !strings.Contains(out, "foo bar") {
		t.Errorf("unexpected output: %s", out)
	}
}
