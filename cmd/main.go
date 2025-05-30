package main

import (
	"flag"

	"github.com/aandrku/portfolio-v2/pkg/handler/site"
	"github.com/aandrku/portfolio-v2/pkg/services/auth"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	tls := flag.Bool("tls", false, "Usage: -tls")
	port := flag.String("port", "3000", "Usage: -port=3000")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	if err = auth.Refresh(); err != nil {
		panic(err)
	}

	e := echo.New()

	site.Register(e)

	if *tls {
		e.Logger.Fatal(e.StartTLS(":"+*port, "tls/certFile", "tls/keyFile"))
	} else {
		e.Logger.Fatal(e.Start(":" + *port))
	}
}
