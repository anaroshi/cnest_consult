package main

import (
	"cnest_consult/univ/app"
	"log"
	"net/http"
	"os"

	"github.com/urfave/negroni"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	
	mux := app.MakeHandler()
	n := negroni.Classic()
	n.UseHandler(mux)

	log.Println("Started App")
	
	log.Fatalln(http.ListenAndServe(":"+port, n))
}