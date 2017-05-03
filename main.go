package main

import (
	"log"
	// "pgu_status/queue"
	"fmt"
	"pgu_status/sx"
	"pgu_status/types"
)

func loadSettings() {
}

// /////////////////// MAIN ///////////////
func main() {
	log.Printf("[INFO] Starting App...")
	log.Printf("[INFO] Load App settings")
	// loadSettings()
	log.Printf("[INFO] End load settings")

	// finder := CreateCaseFinder()
	// defer finder.Close()

	// sxEndpoint := "http://1.99.30.38:8080/sxcs/outgoingRequests"
	sxEndpoint := "http://localhost:8080/sxcs/outgoingRequests"
	sxService := sx.NewSXService(sxEndpoint)

	sxService.ChangePguCaseStatus(types.MakePguStatusMsgStub())

	log.Printf("[INFO] Start listening queue")
	// queue.StartListeningQueue()

	// if deal, err := finder.Find("175426039"); err == nil {
	// 	log.Printf("[INFO] Find CASE: %v", deal)
	// } else {
	// 	log.Printf("[ERROR] SX DB Error: %v", err)
	// }
}

// CreateCaseFinder create
func CreateCaseFinder() *sx.TaskFinder {
	connString := "sx_own01/sx@1.99.17.68:1521/AIXPFMS"
	finder, err := sx.NewTaskFinder(connString)
	if err != nil {
		errMsg := fmt.Sprintf("[ERROR] Can't connect to SX DB. Error: %v", err)
		log.Print(errMsg)
		panic(errMsg)
	}
	return finder
}
