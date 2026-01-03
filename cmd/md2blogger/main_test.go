package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestReadFromStdin(t *testing.T) {
	markdown := "# Test Header"
	stdin := bytes.NewBufferString(markdown)
	var stdout bytes.Buffer

	err := processMarkdown(stdin, &stdout, "")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := stdout.String()
	if !strings.Contains(output, "<h1") {
		t.Errorf("expected HTML header in output, got: %q", output)
	}
}

func TestReadFromFile(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "test*.md")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	markdown := "# File Test"
	if _, err := tmpfile.WriteString(markdown); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	var stdout bytes.Buffer
	err = processMarkdown(nil, &stdout, tmpfile.Name())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := stdout.String()
	if !strings.Contains(output, "<h1") {
		t.Errorf("expected HTML header in output, got: %q", output)
	}
}

func TestHelpFlag(t *testing.T) {
	testCases := []string{"-h", "--help"}

	for _, flag := range testCases {
		t.Run(flag, func(t *testing.T) {
			var stdout bytes.Buffer
			err := processMarkdown(nil, &stdout, flag)

			if err == nil {
				t.Error("expected help flag to return early without error")
			}

			if err.Error() != "help requested" {
				t.Errorf("expected 'help requested' error, got: %v", err)
			}
		})
	}
}
