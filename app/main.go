package main

import (
	"os"
	"transwallet/api/core"
	"transwallet/api/utils/logger"
)

func main() {
	// init logger
	logger.InitLogger()
    log := logger.GetLogger()
	log.Info("Initialize Application")
	// init server
	port := os.Getenv("PORT")
	server := core.NewServer(log)
	log.Info(port)
	if err := server.Listen(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

