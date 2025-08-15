package main

import (
	"flag"
	"github.com/aandrku/personal-website/pkg/services/auth"
	"github.com/joho/godotenv"
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

	app := &application{}

	r := app.router()

	if *tls {
		r.Logger.Fatal(r.StartTLS(":"+*port, "tls/certFile", "tls/keyFile"))
	} else {
		r.Logger.Fatal(r.Start(":" + *port))
	}
}
