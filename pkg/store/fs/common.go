package fs

import (
	"encoding/json"
	"fmt"
	"os"
	"template1/pkg/model"
)

const (
	dataDirectory  = "./data/"
	postsDirectory = dataDirectory + "posts/"
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

func createPostFile(post model.Post) error {
	filename := fmt.Sprintf("%s%s.json", postsDirectory, post.Id.String())
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	if err = encoder.Encode(post); err != nil {
		return err
	}

	return nil
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
