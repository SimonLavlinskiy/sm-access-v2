package main

import (
	"fmt"
	"sm-access/config"
	"sm-access/config/webServer"
	"sm-access/src/controllers"
)

// @title           LICARD BASE
// @version         1.0

// @BasePath  /api/v1
func main() {
	config.Init()

	controllers.InitControllers()

	e := webServer.InitServer()

	//server := socketService.StartServer(messageHandler)

	//go monitoringService.UsbMonitoring(server)

	err := e.Run(":8000")
	if err != nil {
		panic(err)
	}
}

func messageHandler(message []byte) {
	fmt.Println(string(message))
}
