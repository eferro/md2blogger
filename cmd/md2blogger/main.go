package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/eduardoferro/md2blogger/converter"
)

const usage = `md2blogger - Convert Markdown to Blogger-compatible HTML

Usage:
  md2blogger [file]           Convert file to HTML
  cat file.md | md2blogger    Convert stdin to HTML
  md2blogger -h, --help       Show this help

Examples:
  md2blogger post.md > output.html
  cat post.md | md2blogger | pbcopy
`

func main() {
	var filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	if err := processMarkdown(os.Stdin, os.Stdout, filename); err != nil {
		if err.Error() == "help requested" {
			fmt.Print(usage)
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func processMarkdown(stdin io.Reader, stdout io.Writer, filename string) error {
	if filename == "-h" || filename == "--help" {
		return errors.New("help requested")
	}

	var input []byte
	var err error

	if filename != "" {
		input, err = os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("reading file: %w", err)
		}
	} else {
		input, err = io.ReadAll(stdin)
		if err != nil {
			return fmt.Errorf("reading stdin: %w", err)
		}
	}

	html, err := converter.Convert(string(input))
	if err != nil {
		return fmt.Errorf("converting markdown: %w", err)
	}

	_, err = stdout.Write([]byte(html))
	if err != nil {
		return fmt.Errorf("writing output: %w", err)
	}

	return nil
}
