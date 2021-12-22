package main

import (
	"fmt"
	"log"
	"net/http"

	routes "github.com/afeldman/Gakutensoku/upload"
	"github.com/afeldman/go-util/env"
)

func main() {

	port := env.GetEnvOrDefault("GAKUTENSOKU_PORT", "2510")

	log.Println(fmt.Sprintf("Listing on Port: :%s", port))

	http.HandleFunc("/", routes.HTTPHandler)

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

}
