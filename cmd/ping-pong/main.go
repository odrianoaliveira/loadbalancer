package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 8001, "port to listen on")
	flag.Parse()

	_ = fmt.Sprintf("unused variable") // introduce an unused variable to trigger gocilint

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received ping request from %s", r.RemoteAddr)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("pong")); err == nil {
			log.Printf("Sent pong response to %s", r.RemoteAddr)
		} else {
			log.Printf("Failed to write response: %v", err)
		}
	})

	add := fmt.Sprintf(":%d", *port)
	if err := http.ListenAndServe(add, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(1)
	}

	os.Exit(0)
}
