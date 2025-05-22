package main

import (
	"template1/pkg/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	handler.Register(e)

	e.Logger.Fatal(e.Start(":3000"))
}
