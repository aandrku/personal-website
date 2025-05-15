package handler

import (
	"template1/pkg/handler/blog"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo) {
	e.GET("/", indexHandler)
	e.GET("/home", getHomeHandler)
	e.GET("/about", getAboutHandler())
	e.GET("/projects", getProjectsHandler())
	e.GET("/links", getLinksHandler())
	e.GET("/contact", getContactHandler)
	e.GET("/delete", deleteHandler)

	// TODO: this must be protected with middleware
	e.GET("/dashboard", getDashboardHandler)

	blogGroup := e.Group("/blog")
	blog.Register(blogGroup)
}
