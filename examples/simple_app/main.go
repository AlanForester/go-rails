package main

import (
	"go-rails/framework/core"
	"log"
)

func main() {
	app := core.NewApplication()
	log.Println("Starting Go-Rails application...")
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
