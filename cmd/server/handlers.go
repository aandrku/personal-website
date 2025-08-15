package main

import (
	"log"
	"net/http"

	"github.com/aandrku/personal-website/pkg/services/about"
	"github.com/aandrku/personal-website/pkg/services/analytics"
	"github.com/aandrku/personal-website/pkg/services/blog"
	"github.com/aandrku/personal-website/pkg/services/project"
	"github.com/aandrku/personal-website/pkg/store/fs"
	"github.com/aandrku/personal-website/pkg/view"
	"github.com/aandrku/personal-website/pkg/view/home"

	"github.com/labstack/echo/v4"
)

// getIndex serves index page to the client.
func getIndex(c echo.Context) error {
	as := analytics.Service{}
	as.IncrementVisits()

	page := home.Index()

	return view.Render(c, http.StatusOK, page)
}

// getAboutWindow serves about window to the client.
func getAboutWindow(c echo.Context) error {
	info, err := about.GetInfo()
	if err != nil {
		log.Fatalf("failed to get about info: %v\n", err)
	}
	component := home.AboutWindow(info)
	return view.Render(c, http.StatusOK, component)
}

func getProject(c echo.Context) error {
	id := c.Param("id")
	s := project.NewManager(fs.Store{})

	p, err := s.FindProject(id)
	if err != nil {
		return err
	}

	page := home.ProjectPage(p)

	return view.Render(c, http.StatusOK, page)
}

// getHomeWindow serves home window to the client.
func getHomeWindow(c echo.Context) error {
	component := home.HomeWindow()
	return view.Render(c, http.StatusOK, component)
}

func getProjectsWindow(c echo.Context) error {
	s := project.NewManager(fs.Store{})
	p, err := s.Projects()
	if err != nil {
		return err
	}

	component := home.ProjectsWindow(p)
	return view.Render(c, http.StatusOK, component)
}

// getDelete serves empty http response to the client.
//
// This handler is used for removal of html elements, while using HTMX.
func getDelete(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// getWindow serves a blog window to the client.
func getWindow(c echo.Context) error {
	service := blog.NewService()
	posts, err := service.Posts()
	if err != nil {
		return err
	}

	component := home.BlogWindow(posts)
	return view.Render(c, http.StatusOK, component)
}

// getPost serves a post to the client.
func getPost(c echo.Context) error {
	id := c.Param("id")
	service := blog.NewService()

	post, err := service.FindPost(id)
	if err != nil {
		return err
	}

	page := home.PostPage(post)

	return view.Render(c, http.StatusOK, page)
}
