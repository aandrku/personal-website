package blog

import "github.com/labstack/echo/v4"

func Register(g *echo.Group) {
	g.GET("", blogHandler)
	g.GET("/post/:id", postHandler)
}
