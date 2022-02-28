package main

import (
	"github.com/ashishjuyal/banking/app"
	"github.com/ashishjuyal/banking/logger"
)

func main() {

	// log.Println("starting our application...")
	logger.Log.Info("starting the application")
	app.Start()
}
