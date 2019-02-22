package main

import (
	"github.com/davide-c/rest-sample-b/app"
	"github.com/davide-c/rest-sample-b/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}

	app.Initialize(config)
	app.Run(":3000")
}
