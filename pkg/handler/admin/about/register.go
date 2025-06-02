package about

import "github.com/labstack/echo/v4"

func Register(g *echo.Group) {
	// forms
	g.GET("/update-name-form", getUpdateNameForm)
	g.GET("/update-avatar-form", getUpdateAvatarForm)
	g.GET("/update-short-desc-form", getUpdateShortDescForm)
	g.GET("/update-description-form", getUpdateDescriptionForm)

	// update endpoints
	g.POST("/name", updateName)
	g.POST("/avatar", updateAvatar)
	g.POST("/short-desc", updateShortDesc)
	g.POST("/description", updateDescripton)
}
