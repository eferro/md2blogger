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

	err := processMarkdown(stdin, &stdout, "", false)

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
	err = processMarkdown(nil, &stdout, tmpfile.Name(), false)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := stdout.String()
	if !strings.Contains(output, "<h1") {
		t.Errorf("expected HTML header in output, got: %q", output)
	}
}

func TestWithIDsFlag(t *testing.T) {
	markdown := "# Test Header"
	stdin := bytes.NewBufferString(markdown)
	var stdout bytes.Buffer

	err := processMarkdown(stdin, &stdout, "", true)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := stdout.String()
	if !strings.Contains(output, `id="test-header"`) {
		t.Errorf("expected header with ID in output, got: %q", output)
	}
}

func TestWithoutIDsFlag(t *testing.T) {
	markdown := "# Test Header"
	stdin := bytes.NewBufferString(markdown)
	var stdout bytes.Buffer

	err := processMarkdown(stdin, &stdout, "", false)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := stdout.String()
	if strings.Contains(output, `id="`) {
		t.Errorf("expected header without ID in output, got: %q", output)
	}
}
