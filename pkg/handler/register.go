package handler

import (
	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo) {
	e.GET("/", indexHandler)
	e.GET("/home", getHomeHandler)
	e.GET("/about", getAboutHandler)
	e.GET("/projects", getProjectsHandler)
	e.GET("/blog", getBlogHandler)
	e.GET("/links", getLinksHandler)
	e.GET("/contact", getContactHandler)
	e.GET("/delete", deleteHandler)
}
