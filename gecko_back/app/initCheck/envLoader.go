package initCheck

import (
	"github.com/joho/godotenv"
)

// EnvLoader loads the environment variables globallfy
func EnvLoader() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

}
