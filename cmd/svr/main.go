package main

import (
	"net/http"

	"github.com/gazmirmazari/config"
	"github.com/gazmirmazari/kafka_processor/internal/routes"

	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
)

const (
	configPath = "config.yaml"
)

func main() {
	defer quit()

	// Get the config first

	appConfig := config.New(configPath)

	// Set logging level
	// logging.SetLevel(log.Level(appConfig.Logging.Level))

	// Initialize the router
	router := routes.Handler.InitRouter() // Pass the required argument

	// Use gzip for the handler
	gzippedRouter := handlers.CompressHandler(router)

	// Serve the router
	log.Fatal(http.ListenAndServe(appConfig.Port, gzippedRouter))
}

func quit() {
	if r := recover(); r != nil {
		log.Errorf("I panicked and am quitting: %v", r)
		log.Error("I should be alerting someone...")
	}
}
