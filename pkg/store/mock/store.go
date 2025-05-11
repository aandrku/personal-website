package mock

import (
	"errors"
	"template1/pkg/model"

	"github.com/google/uuid"
)

var store *Store

func GetStore() *Store {
	if store == nil {
		store = NewStore()
	}
	return store
}

func NewStore() *Store {
	post1 := &model.Post{
		Id:    uuid.New(),
		Title: "Exploration of tmux",
		Content: `## Title 
- list item 1
- list item 2
		`,
	}
	post2 := &model.Post{
		Id:    uuid.New(),
		Title: "Why I love neovim",
		Content: `
		## Title 
		- list item 1
		- list item 2
		`,
	}
	post3 := &model.Post{
		Id:    uuid.New(),
		Title: "My opinion on the Go programming language",
		Content: `
		## Title 
		- list item 1
		- list item 2
		`,
	}

	store := &Store{
		posts: []*model.Post{post1, post2, post3},
	}

	return store
}

type Store struct {
	posts []*model.Post
}

func (s *Store) GetPosts() []*model.Post {
	return s.posts
}

func (s *Store) GetPostsWithoutContent() []*model.Post {
	return s.posts
}

func (s *Store) FindPost(id string) (*model.Post, error) {
	for _, v := range s.posts {
		if v.Id.String() == id {
			return v, nil
		}
	}
	return nil, errors.New("post not found")
}
