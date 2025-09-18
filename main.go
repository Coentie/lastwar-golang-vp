package main

import (
	"lastwar/notifier/hdadb"
	"lastwar/notifier/lastwar"
	"lastwar/notifier/states"
	"log"
	"os"
	"time"
)

func main() {
	Bootstrap()

	log.Println("----- Application bootstrapped -----")
	log.Println(`----- Checking if conquerer state -----`)

	if lastwar.ConqueredState() {
		log.Println(`----- Oof you got powned, good luck buddy. Conquered state detected -----`)
		_ = os.Setenv("APPLICATION_SATE", states.CONQUERED)
		_ = hdadb.SwipeBottom()
	} else if lastwar.ConquererState() {
		log.Println(`----- Aaay you got them! Conquerer state detected -----`)
		_ = os.Setenv("APPLICATION_SATE", states.CONQUERER)
		_ = hdadb.SwipeBottom()
	} else {
		log.Println(`----- Normal state detected -----`)
		_ = os.Setenv("APPLICATION_SATE", states.NORMAL)
		_ = hdadb.SwipeBottom()
	}

	hdadb.PrintScreen()

	// Run first approval immediately
	log.Println(`----- Running initial loop -----`)
	lastwar.CheckAlert()
	for {
		time.Sleep(1 * time.Minute)
		log.Println("Starting scheduled approval cycle...")
		lastwar.CheckAlert()
	}

}
