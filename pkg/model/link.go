package model

import "github.com/a-h/templ"

// Represents a link to one of my social media accounts.
type Link struct {
	Name    string
	LinkURL string
	Icon    templ.Component
}
