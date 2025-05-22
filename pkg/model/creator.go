package model

import (
	"bytes"
	"io"
	"os"

	"github.com/yuin/goldmark"
)

const descFile = "assets/md/about_description.md"

func NewCreator() (Creator, error) {
	c := Creator{
		Name:      "Andrii Sozonik",
		AvatarURL: "/uploads/avatar.png",
		DescShort: "self-taught developer",
	}

	f, err := os.Open(descFile)
	if err != nil {
		return c, err
	}

	descSrc, err := io.ReadAll(f)
	if err != nil {
		return c, err
	}

	var buf bytes.Buffer
	if err = goldmark.Convert(descSrc, &buf); err != nil {
		return c, err
	}

	c.DescDoc = string(buf.String())

	return c, nil
}

type Creator struct {
	Name      string
	AvatarURL string
	DescShort string
	DescDoc   string
}
