package initialize

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Check() {
	// TODO: will implement logger

	err := godotenv.Load()
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "error!! failed to load .env : %+v", err)
		if err != nil {
			panic(err)
		}

		panic(err)

	}

	ipPath := os.Getenv("LOG_PATH")

	dbIP := os.Getenv("DB_HOST")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")

	gRPCPort := os.Getenv("GPORT")

	var checkList []string
	checkList = append(checkList,
		ipPath,
		dbIP,
		dbPass,
		dbPort,
		dbName,
		dbUser,
		gRPCPort)

	for i, s := range checkList {
		if s == "" {
			_, err := fmt.Fprintf(os.Stderr, "error!! failed to load .env no.%+v : %+v", i, err)
			if err != nil {
				panic(err)
			}

			panic(err)
		}
	}

}
