package store

import (
	"github.com/aandrku/personal-website/pkg/model"
)

type Store interface {
	GetPosts() ([]*model.Post, error)
	FindPost(id string) (*model.Post, error)
	CreatePost(post *model.Post) error
	UpdatePost(post *model.Post) error
	DeletePost(id string) error
}
