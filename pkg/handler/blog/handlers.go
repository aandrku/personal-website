package blog

import (
	"net/http"

	"github.com/aandrku/personal-website/pkg/services/blog"
	"github.com/aandrku/personal-website/pkg/view"
	"github.com/aandrku/personal-website/pkg/view/home"

	"github.com/labstack/echo/v4"
)

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
