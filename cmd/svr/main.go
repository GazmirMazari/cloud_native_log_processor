package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	defer quit()

	//
}

func quit() {
	if r := recover(); r != nil {
		log.Errorf("I panicked and am quitting: %v", r)
		log.Error("I should be alerting someone...")
	}
}
