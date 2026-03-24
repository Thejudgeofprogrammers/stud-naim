package main

import (
	"log"
	"ws-gateway/internal/app"
)

func main() {
	log.Println("starting ws-gateway...")

	app.StartApp()
}
