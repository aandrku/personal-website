package handler

import (
	"log"
	"net/http"
	"template1/pkg/model"
	"template1/pkg/services/project"
	"template1/pkg/view"
	"template1/pkg/view/components"
	"template1/pkg/view/pages"
	"template1/pkg/view/svgs"

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

func getDashboardHandler(c echo.Context) error {
	props := pages.DashboardProps{}

	dashboard := pages.Dashboard(props)

	return view.Render(c, http.StatusOK, dashboard)
}

func getContactHandler(c echo.Context) error {
	component := components.ContactWindow()
	return view.Render(c, http.StatusOK, component)
}

func deleteHandler(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
