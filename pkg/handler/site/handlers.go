package site

import (
	"log"
	"net/http"

	"github.com/aandrku/portfolio-v2/pkg/services/about"
	"github.com/aandrku/portfolio-v2/pkg/services/analytics"
	"github.com/aandrku/portfolio-v2/pkg/services/project"
	"github.com/aandrku/portfolio-v2/pkg/store/fs"
	"github.com/aandrku/portfolio-v2/pkg/view"
	"github.com/aandrku/portfolio-v2/pkg/view/components"
	"github.com/aandrku/portfolio-v2/pkg/view/pages"

	"github.com/labstack/echo/v4"
)

// getIndex serves index page to the client.
func getIndex(c echo.Context) error {
	as := analytics.Service{}
	as.IncrementVisits()

	page := pages.Index()

	return view.Render(c, http.StatusOK, page)
}

// getAboutWindow serves about window to the client.
func getAboutWindow(c echo.Context) error {
	info, err := about.GetInfo()
	if err != nil {
		log.Fatalf("failed to get about info: %v\n", err)
	}
	component := components.AboutWindow(info)
	return view.Render(c, http.StatusOK, component)
}

// getHomeWindow serves home window to the client.
func getHomeWindow(c echo.Context) error {
	component := components.HomeWindow()
	return view.Render(c, http.StatusOK, component)
}

func getProjectsWindow(c echo.Context) error {
	s := project.NewManager(fs.Store{})
	p, err := s.Projects()
	if err != nil {
		return err
	}

	component := components.ProjectsWindow(p)
	return view.Render(c, http.StatusOK, component)
}

// getContactWindow serves contact window to the client.
func getContactWindow(c echo.Context) error {
	component := components.ContactWindow()
	return view.Render(c, http.StatusOK, component)
}

// getDelete serves empty http response to the client.
//
// This handler is used for removal of html elements, while using HTMX.
func getDelete(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
