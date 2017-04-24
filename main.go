package main

import (
  "log"
)

func loadSettings()  {
}

// /////////////////// MAIN ///////////////

func main() {
  log.Printf("[INFO] Starting App...")

  log.Printf("[INFO] Load App settings")
  loadSettings()

  log.Printf("[INFO] Start listening queue")
  startListeningQueue()
  log.Printf("[INFO] End load settings")
}
