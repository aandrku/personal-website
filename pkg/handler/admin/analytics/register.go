package analytics

import "github.com/labstack/echo/v4"

func Register(g *echo.Group) {
	c := newController()
	g.GET("", c.getAnalytics)
}
