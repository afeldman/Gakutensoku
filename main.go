package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

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
		log.Println("INFO: loading .env file failed")
	}

	server := get_env_or_default("SERVER", "127.0.0.1")
	port := get_env_or_default("PORT", "2510")
	wait_t, err := strconv.Atoi(
		*get_env_or_default("WAIT", "30"))

	wait := time.Duration(wait_t)

	r := srv.SetMux()

	srv := &http.Server{
		Handler:      r,
		Addr:         *server + ":" + *port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

}
