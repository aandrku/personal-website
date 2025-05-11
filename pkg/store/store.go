package store

import "template1/pkg/model"

type Store interface {
	GetPosts() ([]*model.Post, error)
	FindPost(id string) (*model.Post, error)
}
