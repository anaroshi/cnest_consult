package main

import (
	"cnest_consult/app"
	"log"
	"net/http"
	"os"

	"github.com/urfave/negroni"
)

func main() {
	mux := app.MakeHandler()
	n := negroni.Classic()
	n.UseHandler(mux)

	log.Println("Started App")
	port := os.Getenv("PORT")
	log.Fatalln(http.ListenAndServe(":"+port, n))
}