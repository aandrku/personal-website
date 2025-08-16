package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aandrku/personal-website/pkg/services/markdown"
	"github.com/aandrku/personal-website/pkg/view/home"
	"gopkg.in/yaml.v3"
)

const (
	Green = "\033[0;32m"
	Red   = "\033[0;31m"
	Reset = "\033[0m"
)

const publicDir = "./public/"

func main() {
	err := os.RemoveAll("./public")
	if err != nil {
	}

	log.Println(Green, "./public succesfully removed!", Reset)

	err = os.Mkdir("public", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("%s failed to create new ./public: %v %s", Red, err, Reset)
	}
	log.Println(Green, "./public succesfully created!", Reset)

	err = generateAbout()
	if err != nil {
		log.Fatalf("%s failed to generate about page: %v %s", Red, err, Reset)
	}
	log.Println(Green, "about page was succesfully generated!", Reset)

}

func generateAbout() error {
	type frontMatter struct {
		ImgURL string `yaml:"image_url"`
	}

	data, err := os.ReadFile("./content/about/about.md")
	if err != nil {
		return err
	}

	yml, md, err := markdown.ExtractYAML(string(data))
	if err != nil {
		return err
	}

	var fm frontMatter
	err = yaml.Unmarshal([]byte(yml), &fm)
	if err != nil {
		return fmt.Errorf("failed to unmarshal front matter")
	}

	html, err := markdown.ToHTML(string(md))
	if err != nil {
		return err
	}

	tmpl := home.AboutWindow(html, fm.ImgURL)

	f, err := os.OpenFile(publicDir+"about.html", os.O_RDWR|os.O_CREATE, 0750)
	if err != nil {
		return err
	}

	err = tmpl.Render(context.Background(), f)
	if err != nil {
		return err
	}

	return nil
}

func generateBlog() error {

	return nil
}

func generateProjects() error {

	return nil
}

func generateMisc() error {

	return nil
}
