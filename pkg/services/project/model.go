package project

import (
	"strings"

	"github.com/aandrku/portfolio-v2/pkg/services/markdown"
	"github.com/google/uuid"
)

type Project struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	ShortDesc    string    `json:"shortDesc"`
	Description  string    `json:"description"`
	Technologies []string  `json:"technologies"`
}

func (p Project) DescriptionHTML() string {
	html, _ := markdown.ToHTML(p.Description)
	return html
}

func (p Project) TechnologiesJoin() string {
	return strings.Join(p.Technologies, ", ")
}
