package main

import (
	"github.com/davide-c/rest-sample-b/app"
	"github.com/davide-c/rest-sample-b/config"
	// "log"
)

func main() {
	config := config.GetConfig()

	// log.Print(config.DB.Name)
	// config.DB.LogMode(true)

	app := &app.App{}

	app.Initialize(config)
	app.Run(":3000")
}
