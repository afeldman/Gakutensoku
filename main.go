package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/afeldman/Gakutensoku/server"
	"github.com/afeldman/go-util/env"
)

func main() {

	port := env.GetEnvOrDefault("GAKUTENSOKU_PORT", "2510")

	log.Println(fmt.Sprintf("Listing on Port: :%s", port))

	http.HandleFunc("/", server.HTTPHandler)

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

}
