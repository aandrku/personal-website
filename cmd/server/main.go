package main

import (
	"flag"
)

func main() {
	tls := flag.Bool("tls", false, "Usage: -tls")
	port := flag.String("port", "3000", "Usage: -port=3000")
	flag.Parse()

	app := &application{}

	r := app.router()

	if *tls {
		r.Logger.Fatal(r.StartTLS(":"+*port, "tls/certFile", "tls/keyFile"))
	} else {
		r.Logger.Fatal(r.Start(":" + *port))
	}
}
