package handler

import (
	"net/http"
	"template1/pkg/view"
	"template1/pkg/view/components"
	"template1/pkg/view/pages"

	"github.com/labstack/echo/v4"
)

func indexHandler(c echo.Context) error {
	page := pages.Index()

	return view.Render(c, http.StatusOK, page)
}

func getAboutHandler(c echo.Context) error {
	content := view.Unsafe("<div>Hello, world</div>")

	component := components.AboutWindow(content)

	return view.Render(c, http.StatusOK, component)
}

func getHomeHandler(c echo.Context) error {
	component := components.HomeWindow()
	return view.Render(c, http.StatusOK, component)
}

func getProjectsHandler(c echo.Context) error {
	component := components.ProjectsWindow()
	return view.Render(c, http.StatusOK, component)
}

func getBlogHandler(c echo.Context) error {
	component := components.BlogWindow()
	return view.Render(c, http.StatusOK, component)
}
func getLinksHandler(c echo.Context) error {
	component := components.LinksWindow()
	return view.Render(c, http.StatusOK, component)
}
func getContactHandler(c echo.Context) error {
	component := components.ContactWindow()
	return view.Render(c, http.StatusOK, component)
}

func deleteHandler(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
