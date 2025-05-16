package admin

import "github.com/labstack/echo/v4"

func Register(g *echo.Group) {
	g.GET("/dashboard", getDashboardHandler)
}
