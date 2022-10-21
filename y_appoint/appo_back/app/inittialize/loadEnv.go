package inittialize

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() map[string]string {
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

	envVar := map[string]string{
		"ipPath":   ipPath,
		"dbIP":     dbIP,
		"dbPass":   dbPass,
		"dbPort":   dbPort,
		"dbName":   dbName,
		"dbUser":   dbUser,
		"gRPCPort": gRPCPort,
	}

	return envVar
}
