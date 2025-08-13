package about

import (
	"fmt"
	"net/http"

	"github.com/aandrku/personal-website/pkg/services/about"
	"github.com/aandrku/personal-website/pkg/services/uploads"
	"github.com/aandrku/personal-website/pkg/view"
	"github.com/aandrku/personal-website/pkg/view/admin"
	"github.com/labstack/echo/v4"
)

func getUpdateNameForm(c echo.Context) error {
	form := admin.UpdateNameForm()
	return view.Render(c, http.StatusOK, form)
}

func getUpdateAvatarForm(c echo.Context) error {
	ups, err := uploads.Get()
	if err != nil {
		return err
	}

	form := admin.UpdateAvatarForm(ups)

	return view.Render(c, http.StatusOK, form)
}

func getUpdateShortDescForm(c echo.Context) error {
	form := admin.UpdateShortDescForm()
	return view.Render(c, http.StatusOK, form)
}

func getUpdateDescriptionForm(c echo.Context) error {
	info, err := about.GetInfo()
	if err != nil {
		return err
	}

	html, err := info.DescriptionHTML()
	if err != nil {
		return nil
	}

	form := admin.UpdateAboutDescriptionForm(html, info.Description)
	return view.Render(c, http.StatusOK, form)
}

func updateName(c echo.Context) error {
	name := c.FormValue("name")

	if err := about.UpdateName(name); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func updateAvatar(c echo.Context) error {
	avatar := c.FormValue("avatar")

	url := fmt.Sprintf("/uploads/%s", avatar)

	if err := about.UpdateAvatarURL(url); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func updateShortDesc(c echo.Context) error {
	sd := c.FormValue("short-desc")

	if err := about.UpdateShortDescription(sd); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func updateDescripton(c echo.Context) error {
	d := c.FormValue("markdown")

	about.UpdateDescription(d)

	return c.NoContent(http.StatusOK)
}
