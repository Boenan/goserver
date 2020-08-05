package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// RunWithDefaultServerConfig start a new server with default config
func RunWithDefaultServerConfig(sm *mux.Router) {
	defaultServer := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	Run(defaultServer)
}

// Run starts the server
func Run(server *http.Server) {

	// TODO: configure the server settings with more specific parameters.

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Server failed to start: %v", err)
			return
		}
	}()

	log.Printf("Server is up and running on %v", server.Addr)

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt, os.Kill)

	sig := <-signalChannel
	log.Printf("Received signal to shutdown server gracefully: %v", sig)

	terminalContext, _ := context.WithTimeout(context.Background(), 30*time.Second)

	err := server.Shutdown(terminalContext)
	if err != nil {
		log.Fatalf("Server failed to shutdown: %v", err)
	}

}
