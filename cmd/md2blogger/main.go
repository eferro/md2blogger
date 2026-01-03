package main

import (
	"fmt"
	"io"
	"os"

	"github.com/eduardoferro/md2blogger/converter"
)

const usage = `md2blogger - Convert Markdown to Blogger-compatible HTML

Usage:
  md2blogger [options] [file]  Convert file to HTML
  cat file.md | md2blogger     Convert stdin to HTML
  md2blogger -h, --help        Show this help

Options:
  -ids    Enable heading IDs in the output

Examples:
  md2blogger post.md > output.html
  md2blogger -ids post.md > output.html
  cat post.md | md2blogger | pbcopy
`

func main() {
	var filename string
	var enableIDs bool

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if arg == "-ids" {
			enableIDs = true
		} else if arg == "-h" || arg == "--help" {
			fmt.Print(usage)
			os.Exit(0)
		} else if len(arg) > 0 && arg[0] == '-' {
			fmt.Fprintf(os.Stderr, "Unknown flag: %s\n\n", arg)
			fmt.Print(usage)
			os.Exit(1)
		} else {
			filename = arg
		}
	}

	if err := processMarkdown(os.Stdin, os.Stdout, filename, enableIDs); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func processMarkdown(stdin io.Reader, stdout io.Writer, filename string, enableIDs bool) error {
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

	html, err := converter.Convert(string(input), enableIDs)
	if err != nil {
		return fmt.Errorf("converting markdown: %w", err)
	}

	_, err = stdout.Write([]byte(html))
	if err != nil {
		return fmt.Errorf("writing output: %w", err)
	}

	return nil
}
