package model

import (
	"github.com/google/uuid"
	"time"
)

func NewPost(title string, content string) *Post {
	return &Post{
		Id:        uuid.New(),
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type Post struct {
	Id        uuid.UUID
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	ViewCount int
}
