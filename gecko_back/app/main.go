package main

import "gecko/router"

func main() {
	e, err := router.Router()

	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Fatal(e.Start(":9010"))
}
