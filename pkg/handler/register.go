package handler

import (
	"github.com/aandrku/personal-website/pkg/handler/site"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Register(e *echo.Echo) {
	// static file serving
	e.Static("/assets", "assets")
	e.Static("/uploads", "./data/uploads")

	// user space group
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	site.Register(e)
}
