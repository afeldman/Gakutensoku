package main

import (
	"log"
	"net/http"
	"os"

	"github.com/afeldman/ktrans_server/src/server"
	"github.com/joho/godotenv"
)

func get_env_or_default(name string, or string) *string {
	ret := os.Getenv(name)
	if (len(ret) == 0) || (ret == "") {
		ret = or
	}
	return &ret
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := get_env_or_default("SERVER", "127.0.0.1")
	port := get_env_or_default("PORT", "2510")

	r := srv.SetMux()

	srv := &http.Server{
		Handler: r,
		Addr:    *server + ":" + *port,
	}

	log.Fatal(srv.ListenAndServe())
}
