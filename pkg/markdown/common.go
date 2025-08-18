package markdown

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	gmhtml "github.com/yuin/goldmark/renderer/html" // alias as gmhtml

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html" // alias as chromahtml
)

var md = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		highlighting.NewHighlighting(
			highlighting.WithStyle("dracula"),
			highlighting.WithFormatOptions(
				chromahtml.WithLineNumbers(true),
				chromahtml.WithClasses(false),
			),
		),
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
	goldmark.WithRendererOptions(
		gmhtml.WithUnsafe(), // allow raw HTML in markdown
	),
)

func ToHTML(markdown string) (string, error) {
	var buf bytes.Buffer

	err := md.Convert([]byte(markdown), &buf)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func ExtractYAML(data string) (string, string, error) {
	chunks := strings.Split(string(data), "---\n")
	if len(chunks) < 3 {
		return "", "", fmt.Errorf("incorrect format")
	}

	yml := chunks[1]
	md := strings.Join(chunks[2:], "---")

	return yml, md, nil
}
