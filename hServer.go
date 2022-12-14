package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/httpMS/handlers"
	// "github.com/httpMS/handlers/types"
)

func main() {

	// Http handlers instance creation
	l := log.New(os.Stdout, "monitor-service", log.LstdFlags)

	// Create HTTP handler for GET handlers
	getAllDataHandler := handlers.NewAllData(l)
	getQueryDataHandler := handlers.NewQuery(l)
	getWebsiteUpdater := handlers.NewUpdateList(l)

	// Create a new serveMux for redirecting the request path
	newServerMux := http.NewServeMux()
	newServerMux.Handle("/", getAllDataHandler)
	newServerMux.Handle("/GET", getQueryDataHandler)
	newServerMux.Handle("/POST", getWebsiteUpdater)

	// Creating a new server configuration
	serverConfig := http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      newServerMux,      // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := serverConfig.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	serverConfig.Shutdown(ctx)

}
