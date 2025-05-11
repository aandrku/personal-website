package blog

import (
	"template1/pkg/model"
	"template1/pkg/store"
	"template1/pkg/store/mock"
)

func NewService() *Service {
	return &Service{
		store: mock.GetStore(),
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

func (s *Service) Posts() []*model.Post {
	return s.store.GetPosts()
}

func (s *Service) PostsWithoutContent() []*model.Post {
	return s.store.GetPostsWithoutContent()

}
