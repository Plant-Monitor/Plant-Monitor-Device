package main

import (
	"log"
	"pcs/snapshots"

	"github.com/joho/godotenv"
)

func main() {
	app := &app{}
	app.setup()

	snapshotPublisher := snapshots.GetSnapshotPublisherInstance()
	snapshotPublisher.Run()
}

type app struct {
}

func (app *app) setup() {
	app.loadEnv()
}

func (app *app) loadEnv() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found")
	}
}
