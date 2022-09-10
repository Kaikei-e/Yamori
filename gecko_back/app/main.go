package main

import "gecko/router"

func main() {
	r, err := router.Router()

	if err != nil {
		r.Logger.Fatal(err)
	}

	r.Logger.Fatal(r.Start(":9000"))
}
