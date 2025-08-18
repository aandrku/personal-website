package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (a *application) router() *echo.Echo {
	e := echo.New()
	e.Static("/assets", "assets")
	e.Static("/uploads", "./data/uploads")

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", getIndex)
	e.GET("/about", getAbout)
	e.GET("/projects", getProjects)
	e.GET("/blog", getBlog)
	e.GET("/misc", getMisc)
	e.GET("/projects/:slug", getProject)
	e.GET("/blog/posts/:slug", getPost)
	e.GET("/misc/posts/:slug", getMiscPost)
	e.GET("/delete", getDelete)

	return e
}
