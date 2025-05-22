package blog

import "github.com/labstack/echo/v4"

func Register(g *echo.Group) {
	// window
	g.GET("/window", getWindow)

	// posts
	g.GET("/post/:id", getPost)
}
