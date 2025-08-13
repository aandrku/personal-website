package fs

import (
	"encoding/json"
	"os"

	"github.com/aandrku/personal-website/pkg/model"
	"github.com/aandrku/personal-website/pkg/services/project"
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
	// project list
	var pl []project.Project

	entries, err := os.ReadDir(projectsDirectory)
	if err != nil {
		return pl, err
	}

	for _, e := range entries {
		var p project.Project
		path := projectsDirectory + e.Name()

		f, err := os.Open(path)
		if err != nil {
			return pl, err
		}

		dec := json.NewDecoder(f)

		if err := dec.Decode(&p); err != nil {
			return pl, err
		}
		pl = append(pl, p)

		f.Close()
	}

	return pl, nil
}

func (s Store) FindProject(id string) (project.Project, error) {
	var p project.Project
	path := projectsDirectory + id + ".json"

	f, err := os.Open(path)
	if err != nil {
		return p, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	if err = dec.Decode(&p); err != nil {
		return p, err
	}
	return p, nil
}

func (s Store) CreateProject(project project.Project) error {
	path := projectsDirectory + project.ID.String() + ".json"

	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(f)

	if err := enc.Encode(project); err != nil {
		return err
	}
	return nil
}

func (s Store) UpdateProject(project project.Project) error {
	path := projectsDirectory + project.ID.String() + ".json"

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(f)

	if err := enc.Encode(project); err != nil {
		return err
	}
	return nil
}
func (s Store) DeleteProject(id string) error {
	path := projectsDirectory + id + ".json"

	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}
