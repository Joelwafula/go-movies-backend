package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	var app application
	app.Domain = "www.google.com"

	log.Println("starting application on port:", port)
	//http.HandleFunc("/", Hello)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
