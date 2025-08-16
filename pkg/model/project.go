package model

import "time"

type Project struct {
	Title        string    `yaml:"title"`
	Slug         string    `yaml:"slug"`
	Technologies []string  `yaml:"technologies"`
	UpdatedAt    time.Time `yaml:"updated_at"`
	CreatedAt    time.Time `yaml:"created_at"`
	Content      string
}
