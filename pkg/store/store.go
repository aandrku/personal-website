package store

import (
	"github.com/aandrku/portfolio-v2/pkg/model"
	"github.com/google/uuid"
)

type Store interface {
	GetPosts() ([]*model.Post, error)
	FindPost(id string) (*model.Post, error)
	CreatePost(post *model.Post) error
	UpdatePost(post *model.Post) error
	DeletePost(id uuid.UUID) error
}
