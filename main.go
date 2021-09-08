package main

import (
	"cnest_consult/app"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	mux := app.MakeHandler()
	n := negroni.Classic()
	n.UseHandler(mux)

	log.Println("Started App")
	log.Fatalln(http.ListenAndServe(":3000", n))
}