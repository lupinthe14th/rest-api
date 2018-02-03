package main

import (
	"log"
	"net/http"
)

func main() {
	api, err := NewRoutes()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
