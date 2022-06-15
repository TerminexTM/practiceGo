package main

import (
	"UdemyREST/app"
	"UdemyREST/logger"
)

func main() {
	//log.Println("Starting our application...") //this is good but we would like to structure our logging
	logger.Info("Starting the application")
	app.Start()
}
