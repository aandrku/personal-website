package blog

import (
	"github.com/aandrku/personal-website/pkg/model"
	"github.com/aandrku/personal-website/pkg/store"
	"github.com/aandrku/personal-website/pkg/store/fs"
)

func NewService() *Service {
	return &Service{
		store: fs.Store{},
	}
}

type Service struct {
	store store.Store
}

// FindPost finds a post by its uuid.
func (s *Service) FindPost(id string) (*model.Post, error) {
	post, err := s.store.FindPost(id)
	if err != nil {
		return &model.Post{}, nil
	}

	return post, nil
}

// CreatePost creates a new post in the filesystem.
func (s *Service) CreatePost(post *model.Post) error {
	if err := s.store.CreatePost(post); err != nil {
		return err
	}
	return nil
}

// UpdatePost updates a post.
func (s *Service) UpdatePost(post *model.Post) error {
	if err := s.store.UpdatePost(post); err != nil {
		return err
	}
	return nil
}

// DeletePost delets a post.
func (s *Service) DeletePost(id string) error {
	if err := s.store.DeletePost(id); err != nil {
		return err
	}
	return nil
}

// Posts returns all posts.
func (s *Service) Posts() ([]*model.Post, error) {
	return s.store.GetPosts()
}
