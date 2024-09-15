package main

import (
	"transwallet/api/core"
	"transwallet/api/utils/logger"
)

func main() {
	// init logger
	logger.InitLogger()
    log := logger.GetLogger()
	log.Info("Initialize Application")
	// init server
	server := core.NewServer(log)
	if err := server.Listen(":8000"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

