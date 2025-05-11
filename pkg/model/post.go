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
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ViewCount int       `json:"viewCount"`
}
