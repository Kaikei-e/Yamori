package main

import (
	"appoint/initialize"
	"appoint/repository"
	"fmt"
	"os"
)

func init() {
	initialize.Check()

	db := repository.NewDBConn()
	err := db.Ping()
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "error!! failed to ping db : %+v", err)
		if err != nil {
			panic(err)
		}

		panic(err)
	}

	d := db.Connect()
	defer d.Close()

}

func main() {
	fmt.Println("hello world")

	db := repository.NewDBConn()
	cn := db.Connect()
	defer cn.Close()

}
