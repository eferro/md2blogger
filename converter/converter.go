package converter

import (
	"bytes"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func Convert(markdown string, enableIDs bool) (string, error) {
	parserOptions := []parser.Option{}
	if enableIDs {
		parserOptions = append(parserOptions, parser.WithAutoHeadingID())
	}

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Table,
		),
		goldmark.WithParserOptions(parserOptions...),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		return "", err
	}

	return addBlankLinesBetweenBlocks(buf.String()), nil
}

func addBlankLinesBetweenBlocks(html string) string {
	blockElements := []string{
		"</p>",
		"</h1>", "</h2>", "</h3>", "</h4>", "</h5>", "</h6>",
		"</ul>", "</ol>",
		"</pre>",
		"</blockquote>",
		"</table>",
		"</div>",
	}

	result := html
	for _, element := range blockElements {
		result = strings.ReplaceAll(result, element+"\n", element+"\n\n")
	}

	return result
}
