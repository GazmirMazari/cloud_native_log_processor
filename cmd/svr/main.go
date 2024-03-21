package main

import (
	"log"
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/gazmirmazari/config"
	"github.com/gazmirmazari/config/logging"
	"github.com/gazmirmazari/internal/routes"
	"github.com/gorilla/handlers"
)

const (
	configPath = "config.yaml"
)

func main() {
	defer quit()

	// Get the config first
	appConfig := config.New(configPath)

	// Set logging level
	logging.SetLevel(log.Level(appConfig.Logging.Level))

	// Initialize the router
	router := routes.InitRouter()

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
