package store

import "template1/pkg/model"

type Store interface {
	GetPosts() []*model.Post
	GetPostsWithoutContent() []*model.Post
	FindPost(id string) (*model.Post, error)
}
