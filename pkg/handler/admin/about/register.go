package about

import "github.com/labstack/echo/v4"

func Register(g *echo.Group) {
	// forms
	g.GET("/form/update-name", getUpdateNameForm)
	g.GET("/form/update-avatar", getUpdateAvatarForm)
	g.GET("/form/update-short-desc", getUpdateShortDescForm)
	g.GET("/form/update-description", getUpdateDescriptionForm)

	// update endpoints
	g.POST("/update-name", postUpdateName)
	g.POST("/update-avatar", postUpdateAvatar)
	g.POST("/update-short-desc", postUpdateShortDesc)
	g.POST("/update-description", postUpdateDescripton)
}
