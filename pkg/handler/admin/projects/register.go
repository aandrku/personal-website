package projects

import "github.com/labstack/echo/v4"

func Register(g *echo.Group) {
	c := newController()

	g.GET("/create-form", c.getCreateFrom)
	g.GET("/:id/update-form", c.getUpdateForm)
	g.GET("/:id/delete-form", c.getDeleteForm)

	g.GET("", c.getProjects)
	g.POST("", c.createProject)
	g.PUT("/:id", c.updateProject)
	g.DELETE("/:id", c.deleteProject)

}
