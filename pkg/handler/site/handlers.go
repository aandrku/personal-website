package site

import (
	"log"
	"net/http"

	"github.com/aandrku/portfolio-v2/pkg/model"
	"github.com/aandrku/portfolio-v2/pkg/services/about"
	"github.com/aandrku/portfolio-v2/pkg/services/analytics"
	"github.com/aandrku/portfolio-v2/pkg/services/project"
	"github.com/aandrku/portfolio-v2/pkg/view"
	"github.com/aandrku/portfolio-v2/pkg/view/components"
	"github.com/aandrku/portfolio-v2/pkg/view/pages"
	"github.com/aandrku/portfolio-v2/pkg/view/svgs"

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

// newGetProjectsWindow return a handler that serves
// projects window to the client.
func newGetProjectsWindow() func(echo.Context) error {
	m := project.NewManager()

	m.AddProject(model.Project{
		Name:         "Hangman TUI",
		DemoURL:      "/assets/pics/hangman-demo.gif",
		Technologies: []string{"Go", "ANSI", "GameDev", "Terminal"},
		MoreInfoFile: "/assets/md/hangman.md",
	})

	m.AddProject(model.Project{
		Name:         "Hangman TUI",
		DemoURL:      "/assets/pics/hangman-demo.gif",
		Technologies: []string{"Go", "ANSI"},
		MoreInfoFile: "/assets/md/hangman.md",
	})
	m.AddProject(model.Project{
		Name:         "Hangman TUI",
		DemoURL:      "/assets/pics/hangman-demo.gif",
		Technologies: []string{"Go", "ANSI"},
		MoreInfoFile: "/assets/md/hangman.md",
	})
	m.AddProject(model.Project{
		Name:         "Hangman TUI",
		DemoURL:      "/assets/pics/hangman-demo.gif",
		Technologies: []string{"Go", "ANSI"},
		MoreInfoFile: "/assets/md/hangman.md",
	})
	m.AddProject(model.Project{
		Name:         "Hangman TUI",
		DemoURL:      "/assets/pics/hangman-demo.gif",
		Technologies: []string{"Go", "ANSI"},
		MoreInfoFile: "/assets/md/hangman.md",
	})
	m.AddProject(model.Project{
		Name:         "Hangman TUI",
		DemoURL:      "/assets/pics/hangman-demo.gif",
		Technologies: []string{"Go", "ANSI"},
		MoreInfoFile: "/assets/md/hangman.md",
	})

	return func(c echo.Context) error {
		component := components.ProjectsWindow(m.Projects())
		return view.Render(c, http.StatusOK, component)
	}

}

// newGetLinksWindow return a handler that serves links window to the client.
func newGetLinksWindow() func(echo.Context) error {
	links := []model.Link{
		{"github", "https://github.com/aandrku", svgs.GithubIcon()},
		{"linkedin", "https://www.linkedin.com/in/aandrku/", svgs.LinkdlnIcon()},
	}

	return func(c echo.Context) error {
		component := components.LinksWindow(links)
		return view.Render(c, http.StatusOK, component)
	}
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
