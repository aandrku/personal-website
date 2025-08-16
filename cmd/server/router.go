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
	e.GET("/about/window", getAboutWindow)
	e.GET("/projects/window", getProjectsWindow)
	e.GET("/projects/:slug", getProject)
	e.GET("/delete", getDelete)
	e.GET("/blog/window", getBlogWindow)
	e.GET("/blog/posts/:slug", getPost)
	e.GET("/misc/window", getMiscWindow)
	e.GET("/misc/posts/:slug", getMiscPost)

	return e
}
