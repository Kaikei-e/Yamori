package main

import (
	"gecko/initCheck"
	"gecko/server"
)

func main() {
	// Check if the .env file exists
	// If it does not exist, panic
	initCheck.EnvLoader()
	initCheck.CheckRun()

	server.Server()

}
