package fs

import (
	"github.com/aandrku/portfolio-v2/pkg/model"
	"github.com/google/uuid"
)

type Store struct{}

func (s Store) GetPosts() ([]*model.Post, error) {
	fileNames, err := readPostFileNames()
	if err != nil {
		return nil, err
	}

	posts := make([]*model.Post, 0)

	for _, f := range fileNames {
		post, err := parsePostFromFile(f)
		if err != nil {
			return posts, err
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

func (s Store) StorePost(post model.Post) error {
	return createPostFile(post)
}

func (s Store) FindPost(id string) (*model.Post, error) {
	posts, err := s.GetPosts()
	if err != nil {
		return &model.Post{}, err
	}

	for _, p := range posts {
		if p.Id.String() == id {
			return p, nil
		}
	}

	return &model.Post{}, nil
}

func (s Store) CreatePost(post *model.Post) error {

	return nil
}

func (s Store) UpdatePost(post *model.Post) error {

	return nil
}

func (s Store) DeletePost(id uuid.UUID) error {

	return nil
}
