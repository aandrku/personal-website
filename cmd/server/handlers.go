package main

import (
	"net/http"

	"github.com/aandrku/personal-website/pkg/view"
	"github.com/aandrku/personal-website/pkg/view/home"

	"github.com/labstack/echo/v4"
)

// getIndex serves index page to the client.
func getIndex(c echo.Context) error {
	return c.File("./public/index.html")
}

// getAboutWindow serves about window to the client.
func getAboutWindow(c echo.Context) error {
	return c.File("./public/about.html")
}

func getProject(c echo.Context) error {
	slug := c.Param("slug")
	return c.File("./public/projects/" + slug + ".html")
}

// getHomeWindow serves home window to the client.
func getHomeWindow(c echo.Context) error {
	component := home.HomeWindow()
	return view.Render(c, http.StatusOK, component)
}

func getProjectsWindow(c echo.Context) error {
	return c.File("./public/projects.html")
}

// getDelete serves empty http response to the client.
//
// This handler is used for removal of html elements, while using HTMX.
func getDelete(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// getBlogWindow serves a blog window to the client.
func getBlogWindow(c echo.Context) error {
	return c.File("./public/blog.html")
}

// getPost serves a post to the client.
func getPost(c echo.Context) error {
	slug := c.Param("slug")

	return c.File("./public/blog/" + slug + ".html")
}
func getMiscWindow(c echo.Context) error {
	return c.File("./public/misc.html")
}

func getMiscPost(c echo.Context) error {
	slug := c.Param("slug")

	return c.File("./public/misc/" + slug + ".html")
}
