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

	db, err := initCheck.CreateDBConn()
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()

	//e, err := router.Router()
	//if err != nil {
	//	e.Logger.Fatal(err)
	//}

	server.Server()

	//err = e.Start(":9010")
	//if err != nil {
	//	crossLogging.Logger.Fatal("Failed to start the server")
	//
	//}
}
