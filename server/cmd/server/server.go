package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Golang-labs-ip/Golang-lab3/server/tablets"
)

//HTTPPortNumber ...
type HTTPPortNumber int

//TabletAPIServer ... Configures necessary handlers and starts listening on a configured port.
type TabletAPIServer struct {
	Port HTTPPortNumber
	TabletsHandler tablets.HTTPHandlerFunc
	server *http.Server
}

// Start will set all handlers and start listening.
// If this methods succeeds, it does not return until server is shut down.
// Returned error will never be nil.
func (s *TabletAPIServer) Start() error {
	if s.TabletsHandler == nil {
		return fmt.Errorf("tablets HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/tablets", s.TabletsHandler)
	
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

// Stop will shut down previously started HTTP server
func (s *TabletAPIServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}