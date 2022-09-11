package main

import (
	"gecko/crossLogging"
	"gecko/initCheck"
	"gecko/router"
)

func main() {
	// Check if the .env file exists
	// If it does not exist, panic
	initCheck.EnvLoader()
	db, err := initCheck.CreateDBConn()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e, err := router.Router()
	if err != nil {
		e.Logger.Fatal(err)
	}

	initCheck.CheckRun(e)

	err = e.Start(":9010")
	if err != nil {
		crossLogging.Logger.Fatal()

	}
}
