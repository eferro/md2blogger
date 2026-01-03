# md2blogger

A modern command-line tool to convert Markdown to Blogger-compatible HTML.

## Features

- **Simple CLI**: Read from file or stdin, output to stdout
- **Full Markdown Support**: Headers, lists, bold, italic, links, images
- **Code Blocks**: Syntax highlighting with language detection
- **Tables**: GitHub Flavored Markdown (GFM) table support
- **Clean HTML**: Formatted, readable HTML output ready for Blogger

## Installation

### Quick Install (Recommended)

Requires Go 1.21 or later:

```bash
git clone https://github.com/eduardoferro/md2blogger.git
cd md2blogger
./install.sh
```

### Manual Install

```bash
git clone https://github.com/eduardoferro/md2blogger.git
cd md2blogger
make install
```

This installs the binary to `~/.local/bin/md2blogger` by default.

**Custom installation directory**:
```bash
make install INSTALL_DIR=/usr/local/bin
```

Make sure the installation directory is in your `PATH`.

### Download Binary

Download pre-built binaries from the [releases page](https://github.com/eduardoferro/md2blogger/releases) (coming soon).

## Usage

### Read from file

```bash
md2blogger mypost.md > output.html
```

### Read from stdin

```bash
cat mypost.md | md2blogger > output.html
```

### Pipe directly to clipboard (macOS)

```bash
md2blogger mypost.md | pbcopy
```

### Pipe directly to clipboard (Linux)

```bash
md2blogger mypost.md | xclip -selection clipboard
```

## Example

Given this Markdown file:

```markdown
# My Blog Post

This is **bold** and this is *italic*.

## Code Example

```go
func main() {
    fmt.Println("Hello!")
}
```

| Feature | Supported |
|---------|-----------|
| Tables  | ✓         |
```

Running `md2blogger example.md` produces:

```html
<h1 id="my-blog-post">My Blog Post</h1>

<p>This is <strong>bold</strong> and this is <em>italic</em>.</p>

<h2 id="code-example">Code Example</h2>

<pre><code class="language-go">func main() {
    fmt.Println("Hello!")
}
</code></pre>

<table>
<thead>
<tr>
<th>Feature</th>
<th>Supported</th>
</tr>
</thead>
<tbody>
<tr>
<td>Tables</td>
<td>✓</td>
</tr>
</tbody>
</table>

```

## Development

See [docs/development_guide.md](docs/development_guide.md) for development instructions.

## License

MIT
