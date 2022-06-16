package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"fib"
)

var revision = "latest"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, revision) })
	mux.Handle("/fibonacci/", http.StripPrefix("/fibonacci", Handler()))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
	log.Printf("Starting fib %s :%s\n", revision, port)
	log.Println(server.ListenAndServe())
}

func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/s/", http.StripPrefix("/s", http.HandlerFunc(sequenceHandler)))
	mux.HandleFunc("/", numberHandler)
	return mux
}

func numberHandler(w http.ResponseWriter, r *http.Request) {
	part := r.URL.Path[1:]
	num, err := strconv.Atoi(part)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintln(w, fib.Fibonacci(num))
}

func sequenceHandler(w http.ResponseWriter, r *http.Request) {
	part := r.URL.Path[1:]
	num, err := strconv.Atoi(part)
	if err != nil {
		num = 10
	}
	next := fib.Sequence()
	for i := 0; i < num; i++ {
		fmt.Fprintln(w, next())
	}
}
