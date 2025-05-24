package project

import "github.com/google/uuid"

type Project struct {
	id           uuid.UUID
	Title        string
	ShortDesc    string
	DemoURL      string
	Description  string
	Technologies string
}
