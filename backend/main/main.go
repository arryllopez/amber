// Starts the amber backend process.
package main

import (
	"log"
	"net/http"

	"amber/backend/config"
	"amber/backend/router"
)

func main() {
	cfg := config.Load()

	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: router.New(),
	}

	log.Printf("amber backend listening on %s", cfg.Addr)
	log.Fatal(server.ListenAndServe())
}
