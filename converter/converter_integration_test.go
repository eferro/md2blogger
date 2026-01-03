package converter

import (
	"strings"
	"testing"
)

func TestConvertComplexMarkdown(t *testing.T) {
	markdown := `# Title

Paragraph with **bold** and *italic*.

## Subtitle

- Item 1
- Item 2

### Code

` + "```go\nfunc main() {}\n```" + `

| Col1 | Col2 |
|------|------|
| A    | B    |

[Link](https://example.com)
`

	result, err := Convert(markdown)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedParts := []string{
		"<h1",
		"<h2",
		"<h3",
		"<strong>bold</strong>",
		"<em>italic</em>",
		"<ul>",
		"<li>",
		"<pre><code",
		"<table>",
		"<a href=",
	}

	for _, part := range expectedParts {
		if !strings.Contains(result, part) {
			t.Errorf("expected result to contain %q, got: %q", part, result)
		}
	}
}
