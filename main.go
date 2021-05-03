package main

import (
	"awesomeProject/app"
	"awesomeProject/logger"
)

// @title Testigo Swagger API
// @version 1.0
// @description Swagger API for Golang Testigo project
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email ratpile@gmail.com

func main() {
	logger.Info("Starting the application..")
	app.Start()
}
