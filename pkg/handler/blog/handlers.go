package blog

import (
	"net/http"
	"template1/pkg/services/blog"
	"template1/pkg/view"
	"template1/pkg/view/components"
	"template1/pkg/view/pages"

	"github.com/labstack/echo/v4"
)

func blogHandler(c echo.Context) error {
	service := blog.NewService()

	component := components.BlogWindow(service.PostsWithoutContent())
	return view.Render(c, http.StatusOK, component)
}

func postHandler(c echo.Context) error {
	id := c.Param("id")
	service := blog.NewService()

	post, err := service.FindPost(id)
	if err != nil {
		return err
	}

	page := pages.PostPage(post)

	return view.Render(c, http.StatusOK, page)
}
