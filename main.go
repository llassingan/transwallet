package main

import (
	"transwallet/app"
	"transwallet/utils/logger"
)

func main() {
	logger.InitLogger()
    log := logger.GetLogger()
	log.Info("Initialize Application")
	server := app.NewServer(log)
	if err := server.Listen(":8000"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

