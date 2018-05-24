package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}

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

	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/products", handler).Methods("POST")
	r.HandleFunc("/articles", handler).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    *server + ":" + *port,
	}

	log.Fatal(srv.ListenAndServe())
}
