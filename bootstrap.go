package main

import (
	"github.com/joho/godotenv"
	"lastwar/notifier/hdadb"
	"log"
)

func Bootstrap() {
	log.Println("Bootstrapping application. Loading .env file")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, are you sure its there?")
	}

	log.Println("Setting up logger...")

	StartLog()

	log.Println("Attempting ADB connection...")

	if err := hdadb.Connect(); err != nil {
		log.Fatalf("Initial ADB connection failed: %v", err)
	}
}
