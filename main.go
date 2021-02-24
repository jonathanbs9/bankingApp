package main

import (
	"github.com/jonathanbs9/bankingApp/app"
	"github.com/jonathanbs9/bankingApp/logger"
)

func main() {
	//log.Println("Starting the application ...")
	logger.Info("starting application ... ")
	app.Start()
}
