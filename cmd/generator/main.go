package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aandrku/personal-website/pkg/model"
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

	err = generateBlog()
	if err != nil {
		log.Fatalf("%s failed to generate blog: %v %s", Red, err, Reset)
	}
	log.Println(Green, "blog was succesfully generated!", Reset)

	err = generateProjects()
	if err != nil {
		log.Fatalf("%s failed to generate projects: %v %s", Red, err, Reset)
	}
	log.Println(Green, "projects was succesfully generated!", Reset)

	err = generateMisc()
	if err != nil {
		log.Fatalf("%s failed to generate misc: %v %s", Red, err, Reset)
	}
	log.Println(Green, "misc was succesfully generated!", Reset)

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
	defer f.Close()

	err = tmpl.Render(context.Background(), f)
	if err != nil {
		return err
	}

	tmpl = home.Index(html, fm.ImgURL)
	f, err = os.OpenFile(publicDir+"index.html", os.O_RDWR|os.O_CREATE, 0750)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Render(context.Background(), f)
	if err != nil {
		return err
	}

	return nil
}

func generateBlog() error {
	err := os.Mkdir(publicDir+"blog", 0750)
	if err != nil {
		return err
	}

	files, err := os.ReadDir("./content/blog")
	if err != nil {
		return err
	}

	posts := make([]model.Postt, 0, len(files))
	for _, f := range files {
		var post model.Postt
		data, err := os.ReadFile("./content/blog/" + f.Name())
		if err != nil {
			return err
		}

		yml, md, err := markdown.ExtractYAML(string(data))
		if err != nil {
			return err
		}

		if err = yaml.Unmarshal([]byte(yml), &post); err != nil {
			return err
		}

		if post.Content, err = markdown.ToHTML(md); err != nil {
			return err
		}

		pFile, err := os.OpenFile(publicDir+"blog/"+post.Slug+".html", os.O_RDWR|os.O_CREATE, 0750)
		if err != nil {
			return err
		}
		defer pFile.Close()

		page := home.PostPage(post)

		page.Render(context.Background(), pFile)

		posts = append(posts, post)
	}

	f, err := os.OpenFile(publicDir+"blog.html", os.O_RDWR|os.O_CREATE, 0750)
	if err != nil {
		return err
	}
	defer f.Close()
	page := home.BlogWindow(posts)

	page.Render(context.Background(), f)

	return nil
}

func generateProjects() error {
	err := os.Mkdir(publicDir+"projects", 0750)
	if err != nil {
		return err
	}

	files, err := os.ReadDir("./content/projects")
	if err != nil {
		return err
	}

	projects := make([]model.Project, 0, len(files))
	for _, f := range files {
		var project model.Project
		data, err := os.ReadFile("./content/projects/" + f.Name())
		if err != nil {
			return err
		}

		yml, md, err := markdown.ExtractYAML(string(data))
		if err != nil {
			return err
		}

		if err = yaml.Unmarshal([]byte(yml), &project); err != nil {
			return err
		}

		if project.Content, err = markdown.ToHTML(md); err != nil {
			return err
		}

		pFile, err := os.OpenFile(publicDir+"projects/"+project.Slug+".html", os.O_RDWR|os.O_CREATE, 0750)
		if err != nil {
			return err
		}
		defer pFile.Close()

		page := home.ProjectPage(project)

		page.Render(context.Background(), pFile)

		projects = append(projects, project)
	}

	f, err := os.OpenFile(publicDir+"projects.html", os.O_RDWR|os.O_CREATE, 0750)
	if err != nil {
		return err
	}
	defer f.Close()
	page := home.ProjectsWindow(projects)

	page.Render(context.Background(), f)

	return nil
}

func generateMisc() error {
	err := os.Mkdir(publicDir+"misc", 0750)
	if err != nil {
		return err
	}

	files, err := os.ReadDir("./content/misc")
	if err != nil {
		return err
	}

	posts := make([]model.Postt, 0, len(files))
	for _, f := range files {
		var post model.Postt
		data, err := os.ReadFile("./content/misc/" + f.Name())
		if err != nil {
			return err
		}

		yml, md, err := markdown.ExtractYAML(string(data))
		if err != nil {
			return err
		}

		if err = yaml.Unmarshal([]byte(yml), &post); err != nil {
			return err
		}

		if post.Content, err = markdown.ToHTML(md); err != nil {
			return err
		}

		pFile, err := os.OpenFile(publicDir+"misc/"+post.Slug+".html", os.O_RDWR|os.O_CREATE, 0750)
		if err != nil {
			return err
		}
		defer pFile.Close()

		page := home.PostPage(post)

		page.Render(context.Background(), pFile)

		posts = append(posts, post)
	}

	f, err := os.OpenFile(publicDir+"misc.html", os.O_RDWR|os.O_CREATE, 0750)
	if err != nil {
		return err
	}
	defer f.Close()
	page := home.MiscWindow(posts)

	page.Render(context.Background(), f)

	return nil
}
