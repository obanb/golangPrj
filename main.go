package main

import (
	"awesomeProject/app"
	"awesomeProject/logger"
)

func main() {
	logger.Info("Starting the application..")
	app.Start()
}
