package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joerdav/webserver/api/handlers/helloworld"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	r := http.NewServeMux()
	r.Handle("/hello", helloworld.Handler{})
	log.Println("Listening on port 4000")
	err := http.ListenAndServe("localhost:4000", r)
	if err != nil {
		return fmt.Errorf("Error when serving: %w", err)
	}
	return nil
}
