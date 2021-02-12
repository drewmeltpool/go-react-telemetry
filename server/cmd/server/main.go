package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/Golang-labs-ip/Golang-lab3/server/db"
	_ "github.com/go-sql-driver/mysql"
)

var httpPortNumber = flag.Int("p", 8080, "HTTP port number")

// NewDbConnection ...
func NewDbConnection() (*sql.DB, error) {
	conn := &db.Connection{
		DbName:   "test_hotel",
		User:     "testuser",
		Password: "12345",
	}
	return conn.Open()
}

func main() {
	// Parse command line arguments. Port number may be defined with "-p" flag.
	flag.Parse()

	// Create the server.
	if server, err := ComposeAPIServer(HTTPPortNumber(*httpPortNumber)); err == nil {
		// Start it.
		go func() {
			log.Println("Starting tablets server...")

			err := server.Start()
			if err == http.ErrServerClosed {
				log.Printf("HTTP server stopped")
			} else {
				log.Fatalf("Cannot start HTTP server: %s", err)
			}
		}()

		// Wait for Ctrl-C signal.
		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt)
		<-sigChannel

		if err := server.Stop(); err != nil && err != http.ErrServerClosed {
			log.Printf("Error stopping the server: %s", err)
		}
	} else {
		log.Fatalf("Cannot initialize tablets server: %s", err)
	}
}
