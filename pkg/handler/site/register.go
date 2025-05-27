package site

import (
	"github.com/aandrku/portfolio-v2/pkg/handler/admin"
	"github.com/aandrku/portfolio-v2/pkg/handler/blog"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo) {
	// static file serving
	e.Static("/assets", "assets")
	e.Static("/uploads", "./data/uploads")

	// index page
	e.GET("/", getIndex)

	// home window
	e.GET("/home-window", getHomeWindow)

	// about window
	e.GET("/about-window", getAboutWindow)

	// projects window
	e.GET("/projects-window", getProjectsWindow)

	// contact window
	e.GET("/contact-window", getContactWindow)

	// delete endpoint, used to delete elements from a webpage using HTMX
	e.GET("/delete", getDelete)

	// blog group
	blog.Register(e.Group("/blog"))

	// admin group
	// TODO: protect with authentication
	admin.Register(e.Group("/admin"))
}
