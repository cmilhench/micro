package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"consignments"
	"consignments/cmd/http/handler"
	"consignments/repository/memory"
)

var revision = "latest"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	service := &consignments.Service{Repo: memory.New()}

	mux := http.NewServeMux()
	mux.Handle("/v", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, revision) }))
	mux.Handle("/", handler.ConsignmentsHandler(service))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	// listen for the context to be canceled in a goroutine
	go func() {
		// block until the context is canceled
		<-ctx.Done()
		// gracefully shutdown the server
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("failed to shutdown server: %v\n", err)
		}
	}()

	log.Printf("Starting consignments %s :%s\n", revision, port)
	log.Printf("%s\n", server.ListenAndServe())
}
