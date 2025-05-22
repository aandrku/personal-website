package blog

import (
	"github.com/aandrku/portfolio-v2/pkg/model"
	"github.com/aandrku/portfolio-v2/pkg/store"
	"github.com/aandrku/portfolio-v2/pkg/store/fs"
)

func NewService() *Service {
	return &Service{
		store: fs.Store{},
	}
}

type Service struct {
	store store.Store
}

// FindPost should find a post by using post's uuid
func (s *Service) FindPost(id string) (*model.Post, error) {
	post, err := s.store.FindPost(id)
	if err != nil {
		return &model.Post{}, nil
	}

	return post, nil
}

func (s *Service) Posts() ([]*model.Post, error) {
	return s.store.GetPosts()
}
