package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"fac/cmd/http/handler"
)

var revision = "latest"

func main() {
	handler := &handler.FactorialHandler{}

	mux := http.NewServeMux()
	mux.HandleFunc("/v", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, revision) })
	mux.Handle("/factorial/", http.StripPrefix("/factorial", handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
	log.Printf("Starting fac %s :%s\n", revision, port)
	log.Println(server.ListenAndServe())
}
