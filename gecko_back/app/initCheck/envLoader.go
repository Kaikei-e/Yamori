package initCheck

import "github.com/joho/godotenv"

func EnvLoader() error {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return nil
}
