package main

import (
	"transwallet/app"
)

func main() {

	server := app.NewServer()
	server.Listen("localhost:8000")
}

