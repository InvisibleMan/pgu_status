package main

import (
	"log"
	// "pgu_status/queue"
)

func loadSettings() {
}

// /////////////////// MAIN ///////////////

func main() {
	log.Printf("[INFO] Starting App...")

	log.Printf("[INFO] Load App settings")
	loadSettings()

	log.Printf("[INFO] Start listening queue")
	// queue.StartListeningQueue()
	log.Printf("[INFO] End load settings")
}
