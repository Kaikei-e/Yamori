package main

import (
	"gecko/crossLogging"
	"gecko/initCheck"
	"gecko/router"
)

func main() {
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
