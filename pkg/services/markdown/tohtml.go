package markdown

import (
	"bytes"

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
