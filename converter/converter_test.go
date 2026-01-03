package converter

import (
	"strings"
	"testing"
)

func TestConvertSimpleMarkdownToHTML(t *testing.T) {
	markdown := "# Hello World"
	expected := `<h1 id="hello-world">Hello World</h1>`

	result, err := Convert(markdown)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if strings.TrimSpace(result) != strings.TrimSpace(expected) {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestParagraphsHaveBlankLineBetween(t *testing.T) {
	markdown := `First paragraph.

Second paragraph.`

	result, err := Convert(markdown)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "<p>First paragraph.</p>\n\n<p>Second paragraph.</p>\n\n"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
