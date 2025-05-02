package handler

import (
	"log"
	"net/http"
	"template1/pkg/model"
	"template1/pkg/view"
	"template1/pkg/view/components"
	"template1/pkg/view/pages"

	"github.com/labstack/echo/v4"
)

func indexHandler(c echo.Context) error {
	page := pages.Index()

	return view.Render(c, http.StatusOK, page)
}

func getAboutHandler() func(echo.Context) error {
	creator, err := model.NewCreator()
	if err != nil {
		log.Fatalf("failed to create creator model: %v\n", err)
	}
	component := components.AboutWindow(creator)

	return func(c echo.Context) error {
		return view.Render(c, http.StatusOK, component)
	}
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
