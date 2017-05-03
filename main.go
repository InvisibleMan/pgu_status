package main

import (
	"log"
	"pgu_status/config"
	"pgu_status/ppot"
	"pgu_status/queue"
	"pgu_status/sx"
)

// /////////////////// MAIN ///////////////
func main() {
	log.Printf("[INFO] Starting App...")

	sxFinder := sx.NewTaskFinder(config.GetString("sx.connString"))
	defer sxFinder.Close()

	sxService := sx.NewSXService(config.GetString("sx.endpoint"))
	msgParser := ppot.NewResultParser()
	listerner := queue.NewListener(config.GetString("ampq.connStr"), config.GetString("ampq.queue"), config.GetString("ampq.errorQueue"))
	defer listerner.Close()

	log.Printf("[INFO] Start listening queue")
	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	listerner.Start(msgParser, sxFinder, sxService)
}
