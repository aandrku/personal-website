package markdown

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/yuin/goldmark"
)

func ToHTML(markdown string) (string, error) {
	var buf bytes.Buffer

	err := goldmark.Convert([]byte(markdown), &buf)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func ExtractYAML(data string) (string, string, error) {
	chunks := strings.Split(string(data), "---\n")
	if len(chunks) != 3 {
		return "", "", fmt.Errorf("incorrect format")
	}

	yml := chunks[1]
	md := chunks[2]

	return yml, md, nil
}
