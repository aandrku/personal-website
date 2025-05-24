package fs

import (
	"encoding/json"
	"os"

	"github.com/aandrku/portfolio-v2/pkg/model"
	"github.com/aandrku/portfolio-v2/pkg/services/project"
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
	path := postsDirectory + post.Filename()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	if err = enc.Encode(post); err != nil {
		return err
	}

	return nil
}

func (s Store) UpdatePost(post *model.Post) error {
	path := postsDirectory + post.Filename()
	f, err := os.OpenFile(path, os.O_TRUNC|os.O_WRONLY, 0)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(f)
	if err = enc.Encode(post); err != nil {
		return err
	}

	return nil
}

func (s Store) DeletePost(id string) error {
	path := postsDirectory + id + ".json"
	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

func (s Store) Projects() ([]project.Project, error) {

	return []project.Project{}, nil
}

func (s Store) FindProject(id string) (project.Project, error) {

	return project.Project{}, nil
}

func (s Store) CreateProject(project project.Project) error {

	return nil
}

func (s Store) UpdateProject(project project.Project) error {

	return nil
}
func (s Store) DeleteProject(project project.Project) error {

	return nil
}
