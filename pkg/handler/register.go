package handler

import (
	"template1/pkg/handler/admin"
	"template1/pkg/handler/blog"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo) {
	// static file serving
	e.Static("/assets", "assets")
	e.Static("/uploads", "./data/uploads")

	e.GET("/", indexHandler)
	e.GET("/home", getHomeHandler)
	e.GET("/about", getAboutHandler())
	e.GET("/projects", getProjectsHandler())
	e.GET("/links", getLinksHandler())
	e.GET("/contact", getContactHandler)
	e.GET("/delete", deleteHandler)

	blogGroup := e.Group("/blog")
	blog.Register(blogGroup)

	// TODO: need to protect this with middleware lateer
	adminGroup := e.Group("/admin")
	admin.Register(adminGroup)
}
