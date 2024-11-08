package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"fib"
)

// This is the entry point for our Microservice, In this case, it will
// bootstrap the HTTP server, listening to port 8080 and route requests to the
// relevant business logic components.

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, fib.ServiceName()) })
	mux.Handle("/fib/", http.StripPrefix("/fib", Handler()))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
	slog.Info(fmt.Sprintf("Starting %s :%s\n", fib.ServiceName(), port))
	if err := server.ListenAndServe(); err != nil {
		slog.Error(fmt.Sprintf("%v", err))
	}
}

func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /n/{num}", numberHandler)
	mux.HandleFunc("GET /s/{num...}", sequenceHandler)
	return mux
}

func numberHandler(w http.ResponseWriter, r *http.Request) {
	val := r.PathValue("num")
	if val == "" {
		http.NotFound(w, r)
		return
	}
	num, err := strconv.Atoi(val)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, fib.Fibonacci(num))
}

func sequenceHandler(w http.ResponseWriter, r *http.Request) {
	val := r.PathValue("num")
	num, err := strconv.Atoi(val)
	if err != nil {
		num = 10
	}
	next := fib.Sequence()
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	for i := 0; i < num; i++ {
		fmt.Fprintln(w, next())
	}
}
