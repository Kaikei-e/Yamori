package initCheck

import (
	"github.com/joho/godotenv"
)

func EnvLoader() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

}
