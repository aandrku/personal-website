package main

import (
	"github.com/aandrku/portfolio-v2/pkg/handler/site"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	site.Register(e)

	e.Logger.Fatal(e.Start(":3000"))
}
