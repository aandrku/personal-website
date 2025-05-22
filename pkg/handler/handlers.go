package handler

import (
	"github.com/aandrku/portfolio-v2/pkg/model"
	"github.com/aandrku/portfolio-v2/pkg/services/analytics"
	"github.com/aandrku/portfolio-v2/pkg/services/project"
	"github.com/aandrku/portfolio-v2/pkg/view"
	"github.com/aandrku/portfolio-v2/pkg/view/components"
	"github.com/aandrku/portfolio-v2/pkg/view/pages"
	"github.com/aandrku/portfolio-v2/pkg/view/svgs"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func indexHandler(c echo.Context) error {
	as := analytics.Service{}
	as.IncrementVisits()

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

func getProjectsHandler() func(echo.Context) error {
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

func getLinksHandler() func(echo.Context) error {
	links := []model.Link{
		{"github", "https://github.com/aandrku", svgs.GithubIcon()},
		{"linkedin", "https://www.linkedin.com/in/aandrku/", svgs.LinkdlnIcon()},
	}

	return func(c echo.Context) error {
		component := components.LinksWindow(links)
		return view.Render(c, http.StatusOK, component)
	}
}

func getContactHandler(c echo.Context) error {
	component := components.ContactWindow()
	return view.Render(c, http.StatusOK, component)
}

func deleteHandler(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
