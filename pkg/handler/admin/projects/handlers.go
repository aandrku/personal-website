package projects

import (
	"net/http"
	"strings"

	"github.com/aandrku/portfolio-v2/pkg/services/project"
	"github.com/aandrku/portfolio-v2/pkg/store/fs"
	"github.com/aandrku/portfolio-v2/pkg/view"
	"github.com/aandrku/portfolio-v2/pkg/view/components/dashboard"
	"github.com/aandrku/portfolio-v2/pkg/view/components/forms"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func newController() *Controller {
	return &Controller{
		service: project.NewManager(fs.Store{}),
	}

}

type Controller struct {
	service *project.Manager
}

func (ct *Controller) getCreateFrom(c echo.Context) error {
	f := forms.CreateProjectForm()

	return view.Render(c, http.StatusOK, f)
}

func (ct *Controller) getUpdateForm(c echo.Context) error {
	id := c.Param("id")

	p, err := ct.service.FindProject(id)
	if err != nil {
		log.Print(err)
		return err
	}

	form := forms.UpdateProjectForm(p)

	return view.Render(c, http.StatusOK, form)
}

func (ct *Controller) getDeleteForm(c echo.Context) error {
	id := c.Param("id")

	p, err := ct.service.FindProject(id)
	if err != nil {
		return err
	}

	form := forms.DeleteProjectForm(p)

	return view.Render(c, http.StatusOK, form)

}

func (ct *Controller) getProjects(c echo.Context) error {
	p, err := ct.service.Projects()
	if err != nil {
		return err
	}

	list := dashboard.ProjectList(p)

	return view.Render(c, http.StatusOK, list)
}

func (ct *Controller) createProject(c echo.Context) error {
	title := c.FormValue("title")
	shortDesc := c.FormValue("short-desc")
	techs := c.FormValue("technologies")
	desc := c.FormValue("markdown")

	technologies := strings.Split(techs, ", ")

	p := project.Project{
		Title:        title,
		ShortDesc:    shortDesc,
		Description:  desc,
		Technologies: technologies,
	}

	if err := ct.service.CreateProject(p); err != nil {
		log.Error(err)
		return err
	}

	c.Response().Header().Add("HX-Trigger", "updateProjects")
	return c.NoContent(http.StatusOK)
}

func (ct *Controller) updateProject(c echo.Context) error {
	title := c.FormValue("title")
	shortDesc := c.FormValue("short-desc")
	techs := c.FormValue("technologies")
	desc := c.FormValue("markdown")

	technologies := strings.Split(techs, ", ")

	p := project.Project{
		Title:        title,
		ShortDesc:    shortDesc,
		Description:  desc,
		Technologies: technologies,
	}

	if err := ct.service.UpdateProject(p); err != nil {
		log.Error(err)
		return err
	}

	c.Response().Header().Add("HX-Trigger", "updateProjects")
	return c.NoContent(http.StatusOK)

}

func (ct *Controller) deleteProject(c echo.Context) error {
	id := c.Param("id")

	if err := ct.service.DeleteProject(id); err != nil {
		log.Print(err)
		return err
	}

	c.Response().Header().Add("HX-Trigger", "updateProjects")
	return c.NoContent(http.StatusOK)

}
