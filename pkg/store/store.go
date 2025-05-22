package store

import "github.com/aandrku/portfolio-v2/pkg/model"

type Store interface {
	GetPosts() ([]*model.Post, error)
	FindPost(id string) (*model.Post, error)
}
