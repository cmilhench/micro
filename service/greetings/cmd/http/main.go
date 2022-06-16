package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"greetings"
)

var revision = "latest"

func main() {
	http.HandleFunc("/v", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, revision) })
	http.HandleFunc("/", rootHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting greetings %s :%s\n", revision, port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "%s\n", greetings.Hello(name))
}
