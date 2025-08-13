package main

import (
	"flag"

	"github.com/aandrku/personal-website/pkg/handler"
	"github.com/aandrku/personal-website/pkg/services/auth"
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

	// authentication OTP must be refreshed at the start
	// or otherwise I would not be able to login into the admin dashboard
	if err = auth.Refresh(); err != nil {
		panic(err)
	}

	e := echo.New()
	handler.Register(e)
	if *tls {
		e.Logger.Fatal(e.StartTLS(":"+*port, "tls/certFile", "tls/keyFile"))
	} else {
		e.Logger.Fatal(e.Start(":" + *port))
	}
}
