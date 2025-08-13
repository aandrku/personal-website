package fs

import (
	"encoding/json"
	"os"

	"github.com/aandrku/personal-website/pkg/model"
)

const (
	dataDirectory     = "./data/"
	postsDirectory    = dataDirectory + "posts/"
	projectsDirectory = dataDirectory + "projects/"
)

func parsePostFromFile(filename string) (model.Post, error) {
	var post model.Post

	f, err := os.Open(postsDirectory + filename)
	if err != nil {
		return post, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	if err = decoder.Decode(&post); err != nil {
		return post, err
	}
	return post, nil
}

func readPostFileNames() ([]string, error) {
	dirs, err := os.ReadDir(postsDirectory)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(dirs))
	for _, d := range dirs {
		names = append(names, d.Name())
	}

	return names, nil
}
