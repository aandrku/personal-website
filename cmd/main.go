package main

import (
	"github.com/aandrku/portfolio-v2/pkg/handler/site"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	site.Register(e)

	e.Logger.Fatal(e.Start(":3000"))
}
